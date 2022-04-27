package com.example.quicgolib.quicgo;

import com.example.quicgolib.util.JsonUtil;

public class Connect {
    public final long connectID;

    Connect(long cid){
        connectID = cid;
    }

    // send message by datagram
    public void send(String message){
         JsonUtil.jsonUtil.fromError(quicgo.Quicgo.sendMessage(connectID,message));
    }

    // get message by datagram
    public String get(){
        return JsonUtil.jsonUtil.fromData(quicgo.Quicgo.receiveMessage(connectID));
    }
}
