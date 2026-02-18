package com.netflix.accountservice.token;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/token")
public class TokenController {

    private final TokenService tokenService;

    public TokenController(TokenService tokenService) {
        this.tokenService = tokenService;
    }

    // POST /token/validate
    @PostMapping("/validate")
    public ResponseEntity<TokenDTO.ValidateResponse> validate(@RequestBody TokenDTO.ValidateRequest request) {
        return ResponseEntity.ok(tokenService.validate(request.getToken()));
    }

    // POST /token/introspect
    @PostMapping("/introspect")
    public ResponseEntity<TokenDTO.IntrospectResponse> introspect(@RequestBody TokenDTO.ValidateRequest request) {
        return ResponseEntity.ok(tokenService.introspect(request.getToken()));
    }
}
