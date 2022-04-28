package com.example.quicgolib.quicgo;

import com.example.quicgolib.util.JsonUtil;

public class ClientConnnect extends Connect{
    ClientConnnect(long cid) {
        super(cid);
    }

    public Stream openStreamSync() {
        return new Stream(JsonUtil.jsonUtil.fromStream(quicgo.Quicgo.openStreamSync(connectID)));
    }
}
