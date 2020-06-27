const LukaText = 1
const LukaImg = 2
const LukaGroup = 1
const LukaSingle = 2


function encodeLukaMsg(from,target,msgType,msgContentType,msg){
    return JSON.stringify(
        {
            CommonField: {
                From: from,
                MsgType: msgType,
                MsgContentType: msgContentType,
                Target: target
            },
            Content: msg
        })
}
