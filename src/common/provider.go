package common

import (
	"container/list"
	"sync"
	"time"
)

//Provider操作
type ProviderOperate interface {
	//根据Session的ID创建Session
	SessionInit(sid string) (Session, error)
	//根据Session的ID读取Session
	SessionRead(sid string) (Session, error)
	//根据Session的ID删除Session
	SessionDestroy(sid string) error
	//对过期的Session进行垃圾回收
	SessionGC(maxLifeTime int64)
}

//Session及其相关操作
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

//Provider
type Provider struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做gc
}

var pder = &Provider{list: list.New()}

//Session存储的结构
type SessionStore struct {
	sid          string                      //session id唯一标示
	timeAccessed time.Time                   //最后访问时间
	value        map[interface{}]interface{} //session里面存储的值
}

//实现Session接口的Set方法
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

//实现Session接口的Set方法
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

//实现Session接口的Delete方法
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

//实现Session接口的SessionID方法
func (st *SessionStore) SessionID() string {
	return st.sid
}

//实现ProviderOperate 的 SessionInit 方法
func (provider *Provider) SessionInit(sid string) (Session, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := provider.list.PushBack(newsess)
	provider.sessions[sid] = element
	return newsess, nil
}

//实现ProviderOperate 的 SessionRead 方法
func (provider *Provider) SessionRead(sid string) (Session, error) {
	if element, ok := provider.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := provider.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

//实现ProviderOperate 的 SessionDestroy 方法
func (provider *Provider) SessionDestroy(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	if element, ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.list.Remove(element)
		return nil
	}
	return nil
}

//实现ProviderOperate 的 SessionGC 方法
func (provider *Provider) SessionGC(maxLifeTime int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	for {
		element := provider.list.Back()
		if element == nil {
			break
		}
		if element.Value.(*SessionStore).timeAccessed.Unix()+maxLifeTime < time.Now().Unix() {
			provider.list.Remove(element)
			delete(provider.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

//自有方法 SessionUpdate
func (provider *Provider) SessionUpdate(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	if element, ok := provider.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		provider.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("memory", *pder)
}
