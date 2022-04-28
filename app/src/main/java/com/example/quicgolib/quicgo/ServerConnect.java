package com.example.quicgolib.quicgo;

import com.example.quicgolib.util.JsonUtil;

public class ServerConnect extends Connect {

    ServerConnect(long cid) {
        super(cid);
    }

    public  Stream acceptStream(){
        return new Stream(JsonUtil.jsonUtil.fromStream(quicgo.Quicgo.acceptStream(connectID)));
    }

}
