package com.coderconsole.tdd.testdata;

import com.coderconsole.tdd.ShoppingCart;

import java.util.ArrayList;
import java.util.List;

public class ShoppingCartData {

    public static List<ShoppingCart.Cart> getMockList() {
        List<ShoppingCart.Cart> carts = new ArrayList<>();
        carts.add(new ShoppingCart.Cart("Onion", 30));
        return carts;
    }
}
