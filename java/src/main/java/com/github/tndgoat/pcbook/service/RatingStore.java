package com.github.tndgoat.pcbook.service;

public interface RatingStore {
    Rating Add(String laptopID, double score);
}