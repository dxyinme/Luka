const LukaText = 1
const LukaImg = 2
const LukaVideo = 3
const LukaGroup = 1
const LukaSingle = 2


function encodeLukaMsg(from,target,msgType,msgContentType,msg){
    var myDate = new Date();
    return JSON.stringify(
        {
            From: from,
            MsgType: msgType,
            MsgTime: myDate.toLocaleString(),
            MsgContentType: msgContentType,
            Target: target,
            Content: msg
        })
}
