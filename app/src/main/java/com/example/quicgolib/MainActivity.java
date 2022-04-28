package com.example.quicgolib;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.text.Editable;
import android.view.View;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

import com.example.quicgolib.quicgo.ClientConnnect;
import com.example.quicgolib.quicgo.Connect;
import com.example.quicgolib.quicgo.QuicGoClientEngine;
import com.example.quicgolib.quicgo.QuicGoServerEngine;
import com.example.quicgolib.quicgo.ServerConnect;
import com.example.quicgolib.quicgo.Stream;
import com.example.quicgolib.util.Logger;

import quicgo.Quicgo;

public class MainActivity extends AppCompatActivity {

    static final String TAG = "MainActivity";

    QuicGoServerEngine serverEngine = new QuicGoServerEngine("127.0.0.1:41420");
    QuicGoClientEngine clientEngine = new QuicGoClientEngine("127.0.0.1:41420");
    ServerConnect serverConn;
    ClientConnnect clientConn;
    Stream clientStream;
    Stream serverStream;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        initButton();

        new Thread(new Runnable() {
            @Override
            public void run() {
                Logger.d(TAG,"server listen to the connection... ");
                serverConn = (ServerConnect) serverEngine.listen().accept();
                Logger.d(TAG,"server accept the connection: connectID "+ serverConn.connectID);

                Logger.d(TAG,"server listen to the stream... ");
                // 只有相应的stream有data才会accept
                serverStream =  serverConn.acceptStream();
                Logger.d(TAG,"server accept the stream: streamID "+ serverStream.streamID);
                while(true) {
                    String message = serverStream.get();
                    Logger.d(TAG,"Server get the streamMessage: "+ message);
                    serverStream.send("server return echo :" + message);
                }
            }
        }).start();

        clientConn = (ClientConnnect) clientEngine.dial();
        Logger.d(TAG,"client build the connection: connectID "+ clientConn.connectID);

        clientStream =  clientConn.openStreamSync();
        Logger.d(TAG,"client build the stream: streamID "+ clientStream.streamID);
    }

    private void initButton(){
        Button streamButton = findViewById(R.id.stream_send_button);
        Button datagramButton = findViewById(R.id.datagram_send_button);

        streamButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                EditText streamMessage = findViewById(R.id.stream_message);
                Editable messageText = streamMessage.getText();


                if(!messageText.equals("")){
                    clientStream.send(messageText.toString());
                    Logger.d(TAG,"Client send the streamMessage: "+ messageText.toString());
                }else{
                    Logger.e(TAG, "Can't send empty message");
                }
                TextView resultText = findViewById(R.id.result_text);
                resultText.setText("Client get message: [" + clientStream.get() + "]");
            }
        });

        datagramButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                EditText streamMessage = findViewById(R.id.datagram_message);
                Editable messageText = streamMessage.getText();


                if(!messageText.equals("")){
                    Logger.d(TAG, "Client send message: "+ messageText.toString());
                    clientConn.send(messageText.toString());
                }else{
                    Logger.e(TAG, "Can't send empty message");
                }
                serverConn.send("server return echo :" + serverConn.get());
                TextView resultText = findViewById(R.id.result_text);
                resultText.setText("Client get message: [" + clientConn.get() + "]");
            }
        });
    }


}