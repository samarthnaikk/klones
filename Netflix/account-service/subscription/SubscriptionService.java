package com.netflix.accountservice.subscription;

import org.springframework.stereotype.Service;

@Service
public class SubscriptionService {

    private final SubscriptionRepository subscriptionRepository;

    public SubscriptionService(SubscriptionRepository subscriptionRepository) {
        this.subscriptionRepository = subscriptionRepository;
    }

    public SubscriptionEntity getSubscription(Long userId) {
        return subscriptionRepository.findByUserId(userId)
                .orElseThrow(() -> new RuntimeException("Subscription not found for user: " + userId));
    }

    public SubscriptionEntity updateSubscription(Long userId, String plan) {
        SubscriptionEntity sub = getSubscription(userId);
        sub.setPlan(plan);
        return subscriptionRepository.save(sub);
    }

    public SubscriptionEntity cancelSubscription(Long userId) {
        SubscriptionEntity sub = getSubscription(userId);
        sub.setStatus("CANCELLED");
        return subscriptionRepository.save(sub);
    }
}
