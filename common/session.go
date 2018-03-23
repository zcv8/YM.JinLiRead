package common

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
	_"log"
)

//用户登录的CookieName
const AuthorizationKey = "AUTHORIZATIONKEY"

//Session管理器
type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    Provider
	maxLifeTime int64
}

var provides = make(map[string]Provider)

//注册Provider
func Register(name string, provider Provider) {
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}

//创建Session管理器
func NewManager(providerName string, cookieName string, maxLifeTime int64) (*SessionManager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &SessionManager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

//根据每次用户访问发现与之关联的Session,没有关联的Session则创建
func (manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request, saveTime int64) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	//没发现cookie或者是cookie的值为空，则创建Session
	if err != nil || cookie.Value == "" {
		sid := GetGuid()
		session, _ = manager.provider.SessionInit(sid, saveTime)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid, saveTime)
	}
	return
}

//读取Session
func (manager *SessionManager) SessionRead(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	cookie, err := r.Cookie(manager.cookieName)
	if err == nil && cookie.Value != "" {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.OverSessionRead(sid)
	}
	return
}

//销毁Session
func (manager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(AuthorizationKey)
	if err == nil && cookie.Value != "" {
		manager.lock.Lock()
		defer manager.lock.Unlock()
		sid, _ := url.QueryUnescape(cookie.Value)
		manager.provider.SessionDestroy(sid)
	}
	return
}

//对过期的Session进行垃圾回收
func (manager *SessionManager) GC() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionGC(manager.maxLifeTime)
	//每隔 maxLifeTime 的时间再次进行回收
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() { manager.GC() })
}
