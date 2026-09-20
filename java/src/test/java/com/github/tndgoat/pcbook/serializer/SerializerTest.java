package com.github.tndgoat.pcbook.serializer;

import com.github.tndgoat.pcbook.pb.Laptop;
import com.github.tndgoat.pcbook.sample.Generator;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import java.io.IOException;

class SerializerTest {

    @Test
    void writeAndReadBinaryFile() throws IOException {
        String binaryFile = "laptop.bin";
        Laptop laptop1 = new Generator().newLaptop();

        Serializer serializer = new Serializer();
        serializer.writeBinaryFile(laptop1, binaryFile);

        Laptop laptop2 = serializer.readBinaryFile(binaryFile);
        Assertions.assertEquals(laptop1, laptop2);
    }
}