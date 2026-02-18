package com.netflix.accountservice.concurrency;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/concurrency")
public class ConcurrencyController {

    private final ConcurrencyService concurrencyService;

    public ConcurrencyController(ConcurrencyService concurrencyService) {
        this.concurrencyService = concurrencyService;
    }

    // GET /concurrency/check
    @GetMapping("/check")
    public ResponseEntity<Boolean> check(@RequestParam Long userId,
                                         @RequestParam String profileId) {
        return ResponseEntity.ok(concurrencyService.check(userId, profileId));
    }

    // POST /concurrency/lock
    @PostMapping("/lock")
    public ResponseEntity<ConcurrencyEntity> lock(@RequestBody ConcurrencyDTO.LockRequest request) {
        return ResponseEntity.ok(concurrencyService.lock(
                request.getUserId(), request.getProfileId(), request.getStreamId()));
    }

    // DELETE /concurrency/release
    @DeleteMapping("/release")
    public ResponseEntity<Void> release(@RequestParam String streamId) {
        concurrencyService.release(streamId);
        return ResponseEntity.noContent().build();
    }
}
