package com.example.quicgolib.quicgo;

import com.example.quicgolib.util.JsonUtil;

import quicgo.Quicgo;

public class QuicGoClientEngine {
    private final String address;

    public QuicGoClientEngine(String addr){
        address = addr;
    }

    public Connect dial(){
        return new ClientConnnect(JsonUtil.jsonUtil.fromConnect(quicgo.Quicgo.dial(address)));
    }
}
