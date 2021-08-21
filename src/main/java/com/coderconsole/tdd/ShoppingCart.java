package com.coderconsole.tdd;

import java.util.ArrayList;
import java.util.List;

public class ShoppingCart {

    List<Cart> carts = new ArrayList<>();

    public List<Cart> getCartList() {
        return carts;
    }

    public List<Cart> addCart(String name, long price) {
        carts.add(new Cart(name, price));
        return carts;
    }

    public static class Cart {
        String name;
        long price;

        public Cart(String name, long price) {
            this.name = name;
            this.price = price;
        }
    }
}
