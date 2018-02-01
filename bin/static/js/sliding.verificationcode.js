 var verifycode = function() {
    var dragContainer = document.getElementById("dragContainer");
    var dragBg = document.getElementById("dragBg");
    var dragText = document.getElementById("dragText");
    var dragHandler = document.getElementById("dragHandler");

    //滑块最大偏移量
    var maxHandlerOffset = 306;
    //是否验证成功的标记
    var isVertifySucc = false;
    initDrag();

    function initDrag() {
        dragText.textContent = "拖动滑块验证";
        dragHandler.addEventListener("mousedown", onDragHandlerMouseDown);
    }

    function onDragHandlerMouseDown() {
        dragHandler.addEventListener("mousemove", onDragHandlerMouseMove);
        dragHandler.addEventListener("mouseup", onDragHandlerMouseUp);
    }

    function onDragHandlerMouseMove() {
        
    }
    function onDragHandlerMouseUp() {
        dragHandler.removeEventListener("mousemove", onDragHandlerMouseMove);
        dragHandler.removeEventListener("mouseup", onDragHandlerMouseUp);
        dragHandler.style.left = 0;
        dragBg.style.width = 0;
    }

    //验证成功
    function verifySucc() {
        isVertifySucc = true;
        dragText.textContent = "验证通过";
        dragText.style.color = "white";
        dragHandler.setAttribute("class", "dragHandlerOkBg");
        dragHandler.removeEventListener("mousedown", onDragHandlerMouseDown);
        dragHandler.removeEventListener("mousemove", onDragHandlerMouseMove);
        dragHandler.removeEventListener("mouseup", onDragHandlerMouseUp);
    };
}