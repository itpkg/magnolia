package com.github.itpkg.magnolia.auth.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

/**
 * Created by bandari on 16-10-5.
 */
@Controller("auth.usersController")
@RequestMapping(path = "/users")
public class UsersController {
    @GetMapping("/sign-in")
    public String getSignIn(Model model) {
        return "users/sign-in";
    }
}
