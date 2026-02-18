package com.netflix.accountservice.subscription;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/subscriptions")
public class SubscriptionController {

    private final SubscriptionService subscriptionService;

    public SubscriptionController(SubscriptionService subscriptionService) {
        this.subscriptionService = subscriptionService;
    }

    // GET /subscriptions/{userId}
    @GetMapping("/{userId}")
    public ResponseEntity<SubscriptionEntity> getSubscription(@PathVariable Long userId) {
        return ResponseEntity.ok(subscriptionService.getSubscription(userId));
    }

    // PATCH /subscriptions/{userId}
    @PatchMapping("/{userId}")
    public ResponseEntity<SubscriptionEntity> updateSubscription(@PathVariable Long userId,
                                                                  @RequestParam String plan) {
        return ResponseEntity.ok(subscriptionService.updateSubscription(userId, plan));
    }

    // POST /subscriptions/{userId}/cancel
    @PostMapping("/{userId}/cancel")
    public ResponseEntity<SubscriptionEntity> cancelSubscription(@PathVariable Long userId) {
        return ResponseEntity.ok(subscriptionService.cancelSubscription(userId));
    }
}
