package com.example.quicgolib.util;

import android.util.Log;

import com.google.gson.Gson;
import com.google.gson.annotations.SerializedName;

public class JsonUtil {

    public final static JsonUtil jsonUtil = new JsonUtil();

    private JsonUtil(){}

    private final Gson gson = new Gson();

    private static String TAG = "QUIC-GO";

    private static class ErrorReturn{
        @SerializedName("error")
        String error;
    }

    private static class ConnectReturn{
        @SerializedName("error")
        String error;

        @SerializedName("connect_id")
        int connectID;
    }

    private static class StreamReturn{
        @SerializedName("error")
        String error;

        @SerializedName("stream_id")
        int streamID;
    }

    private static class DataReturn{
        @SerializedName("error")
        String error;

        @SerializedName("data")
        String data;
    }

    public void fromError(String jsonData){
        String err = gson.fromJson(jsonData,ErrorReturn.class).error;
        if(!err.equals("")){
            Log.e(TAG,err);
            System.exit(1);
        }
    }

    public long fromConnect(String jsonData){
        ConnectReturn connectReturn = gson.fromJson(jsonData,ConnectReturn.class);
        if(!connectReturn.error.equals("")){
            Log.e(TAG,connectReturn.error);
            System.exit(1);
        }
        return connectReturn.connectID;
    }

    public long fromStream(String jsonData){
        StreamReturn streamReturn = gson.fromJson(jsonData,StreamReturn.class);
        if(!streamReturn.error.equals("")){
            Log.e(TAG, streamReturn.error);
            System.exit(1);
        }
        return streamReturn.streamID;
    }

    public String fromData(String jsonData){
        DataReturn dataReturn = gson.fromJson(jsonData, DataReturn.class);
        if(!dataReturn.error.equals("")){
            Log.e(TAG, dataReturn.error);
            System.exit(1);
        }
        return dataReturn.data;
    }

}
