package com.example.server;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.NOT_FOUND)
public class ProjectNotFoundException extends RuntimeException {

    private final String name;

    public ProjectNotFoundException(String name) {
        super(name);
        this.name = name;
    }

    public String getName() {
        return name;
    }
}
