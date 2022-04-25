package com.example.quicgolib;

import androidx.appcompat.app.AppCompatActivity;

import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

import quicgosdk.Quicgosdk;

public class MainActivity extends AppCompatActivity {

    static final String TAG = "MainActivity";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Button button = findViewById(R.id.button);
        Thread thread = new Thread(new Runnable() {
            @Override
            public void run() {
                Log.w( TAG ,Quicgosdk.buildServer("127.0.0.1:41420"));
            }
        });
        thread.start();
        button.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                TextView text = findViewById(R.id.text);
                text.setText(Quicgosdk.buildClient("127.0.0.1:41420","Hello quic-go!!"));
            }
        });
    }
}