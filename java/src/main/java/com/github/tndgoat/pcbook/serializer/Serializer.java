package com.github.tndgoat.pcbook.serializer;

import com.github.tndgoat.pcbook.pb.Laptop;
import com.google.protobuf.util.JsonFormat;

import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.charset.StandardCharsets;

public class Serializer {

    public void writeBinaryFile(Laptop laptop, String filename) throws IOException {
        try (FileOutputStream outStream = new FileOutputStream(filename)) {
            laptop.writeTo(outStream);
        }
    }

    public Laptop readBinaryFile(String filename) throws IOException {
        try (FileInputStream inStream = new FileInputStream(filename)) {
            return Laptop.parseFrom(inStream);
        }
    }

    public void writeJSONFile(Laptop laptop, String filename) throws IOException {
        JsonFormat.Printer printer = JsonFormat.printer()
                .includingDefaultValueFields()
                .preservingProtoFieldNames();

        String jsonString = printer.print(laptop);

        try (FileOutputStream outStream = new FileOutputStream(filename)) {
            outStream.write(jsonString.getBytes(StandardCharsets.UTF_8));
        }
    }

    public static void main(String[] args) throws IOException {
        Serializer serializer = new Serializer();
        Laptop laptop = serializer.readBinaryFile("laptop.bin");
        serializer.writeJSONFile(laptop, "laptop.json");
    }
}