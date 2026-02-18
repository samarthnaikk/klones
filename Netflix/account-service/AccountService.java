package com.netflix.accountservice;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * AccountService is the sole entry point for the Account microservice.
 * It bootstraps all domain modules: user, auth, session, profile,
 * device, subscription, entitlement, concurrency, token, security.
 */
@SpringBootApplication(scanBasePackages = "com.netflix.accountservice")
public class AccountService {
    public static void main(String[] args) {
        SpringApplication.run(AccountService.class, args);
    }
}
