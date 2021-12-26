package com.coderconsole.tdd;

import java.util.ArrayList;
import java.util.List;
import java.util.Optional;

/**
 * Validate the email lookup
 *
 * ["gmail.com", "yahoo.com", "outlook.com"]
 * nitesh.tiwari@in.com
 *
 */
public class EmailLookUp {

    private final List<String> lookUpList;

    public EmailLookUp(){
        lookUpList = new ArrayList<>();
        lookUpList.add("gmail.com");
        lookUpList.add("outlook.com");
        lookUpList.add("yahoo.com");
    }

    public boolean validateEmail(String email) {
        if (email == null || email.isEmpty())
            return false;
        String domainName = getDomainName(email);
        Optional<String> count = lookUp(domainName);
        return count.isPresent();
    }

    private Optional<String> lookUp(String domainName) {
        return lookUpList
                .stream()
                .filter(lookUp -> lookUp.equals(domainName))
                .findFirst();
    }

    public String getDomainName(String email) {
        int index = email.lastIndexOf("@");
        return email.substring(index + 1);
    }



}
