package com.netflix.accountservice.session;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/sessions")
public class SessionController {

    private final SessionService sessionService;

    public SessionController(SessionService sessionService) {
        this.sessionService = sessionService;
    }

    // GET /sessions/active
    @GetMapping("/active")
    public ResponseEntity<List<SessionEntity>> getActiveSessions(@RequestParam Long userId) {
        return ResponseEntity.ok(sessionService.getActiveSessions(userId));
    }

    // DELETE /sessions/{id}
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteSession(@PathVariable String id) {
        sessionService.deleteSession(id);
        return ResponseEntity.noContent().build();
    }
}
