package com.example.quicgolib.quicgo;

import com.example.quicgolib.util.JsonUtil;

public class Stream {
    public final long streamID;

    Stream(long sid){
        streamID = sid;
    }

    public void send(String message){
        JsonUtil.jsonUtil.fromError(quicgo.Quicgo.sendMessage(streamID,message));
    }

    public String get(){
        return JsonUtil.jsonUtil.fromData(quicgo.Quicgo.receiveMessage(streamID));
    }
}
