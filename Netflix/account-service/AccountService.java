package com.netflix.accountservice;

import com.netflix.accountservice.auth.AuthApplication;
import com.netflix.accountservice.user.UserApplication;
import com.netflix.accountservice.profile.ProfileApplication;
import com.netflix.accountservice.device.DeviceApplication;
import com.netflix.accountservice.subscription.SubscriptionApplication;
import com.netflix.accountservice.entitlement.EntitlementApplication;

public class AccountService {
    public static void main(String[] args) {
        // This is a placeholder for integrating all microservices.
        // In a real Spring Boot monorepo, you might use Spring Cloud or modules.
        System.out.println("AccountService root application started.");
        // AuthApplication.main(args);
        // UserApplication.main(args);
        // ProfileApplication.main(args);
        // DeviceApplication.main(args);
        // SubscriptionApplication.main(args);
        // EntitlementApplication.main(args);
    }
}
