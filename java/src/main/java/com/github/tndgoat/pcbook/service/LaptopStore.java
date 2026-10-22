package com.github.tndgoat.pcbook.service;

import com.github.tndgoat.pcbook.pb.Laptop;

public interface LaptopStore {
    void Save(Laptop laptop) throws Exception;
    Laptop Find(String id);
}
