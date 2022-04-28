package com.example.quicgolib.quicgo;

import com.example.quicgolib.util.JsonUtil;
import com.example.quicgolib.util.Logger;

public class QuicGoServerEngine {

    private final String address;

    public QuicGoServerEngine(String addr){
        address = addr;
    }

    public QuicGoServerEngine listen(){
        JsonUtil.jsonUtil.fromError(quicgo.Quicgo.listen(address));
        return this;
    }

    public Connect accept(){
        return new ServerConnect(JsonUtil.jsonUtil.fromConnect(quicgo.Quicgo.accept()));
    }
}
