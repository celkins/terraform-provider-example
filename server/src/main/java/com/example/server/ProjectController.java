package com.example.server;

import org.springframework.web.bind.annotation.*;

import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

import static org.springframework.http.MediaType.APPLICATION_JSON_VALUE;

@RestController
@RequestMapping("projects")
public class ProjectController {

    private final Map<String, Project> projects = new ConcurrentHashMap<>();

    @GetMapping(path = "{name}", produces = APPLICATION_JSON_VALUE)
    public Project getProject(@PathVariable String name) {
        //FIXME Race condition
        if (!this.projects.containsKey(name)) {
            throw new ProjectNotFoundException(name);
        }
        return this.projects.get(name);
    }

    @PostMapping(path = "{name}", consumes = APPLICATION_JSON_VALUE)
    public void createProject(@PathVariable String name, @RequestBody Project project) {
        this.projects.put(name, project);
    }

    @PutMapping(path = "{name}", consumes = APPLICATION_JSON_VALUE)
    public void updateProject(@PathVariable String name, @RequestBody Project project) {
        this.projects.put(name, project);
    }

    @DeleteMapping(path = "{name}")
    public void deleteProject(@PathVariable String name) {
        if (this.projects.remove(name) == null) {
            throw new ProjectNotFoundException(name);
        }
    }
}
