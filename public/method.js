function renderScreen(ctx,cvs) {
    return new Promise(function (res,req) {
        scale = 3

        img = new Image()
        img.src = "/screen"
        img.onload = function () {
            cvs.width=img.width/scale
            cvs.height=img.height/scale
            ctx.drawImage(img,0,0,img.width/scale,img.height/scale)
            res()
        }
    })
}