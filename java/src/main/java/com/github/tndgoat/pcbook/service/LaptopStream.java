package com.github.tndgoat.pcbook.service;

import com.github.tndgoat.pcbook.pb.Laptop;

public interface LaptopStream {
    void Send(Laptop laptop);
}