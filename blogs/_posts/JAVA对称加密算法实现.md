---
title: JAVA对称加密算法实现
mathjax: false
date: 2023-04-04 19:01:46
tags:
categories:
---

# 对称加密算法

> 将数据加密后还可以解密获取原始信息的算法.

想我们的AES算法就是一种十分安全的加密算法就是说 还是不错的哈哈哈.

# 类的设计

* String getSalt  生成固定长度的随机数 用作加密salt
* String Encoding 编码
* String  解码

```java
package demo;

import java.security.SecureRandom;
import java.util.Base64;
import javax.crypto.Cipher;
import javax.crypto.spec.IvParameterSpec;
import javax.crypto.spec.SecretKeySpec;

public class Encryption {
  private static final int SALT_LENGTH = 48;
  
  public static String getSalt() {
    SecureRandom random = new SecureRandom();
    byte[] salt = new byte[SALT_LENGTH];
    random.nextBytes(salt);
    return Base64.getEncoder().encodeToString(salt);
  }
  
  public static String encode(String input, String salt) throws Exception {
    byte[] inputBytes = input.getBytes("UTF-8");
    byte[] saltBytes = Base64.getDecoder().decode(salt);
    byte[] keyBytes = new byte[16];
    byte[] ivBytes = new byte[16];
    System.arraycopy(saltBytes, 0, keyBytes, 0, 16);
    System.arraycopy(saltBytes, 16, ivBytes, 0, 16);
    SecretKeySpec keySpec = new SecretKeySpec(keyBytes, "AES");
    IvParameterSpec ivSpec = new IvParameterSpec(ivBytes);
    Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
    cipher.init(Cipher.ENCRYPT_MODE, keySpec, ivSpec);
    byte[] encryptedBytes = cipher.doFinal(inputBytes);
    return Base64.getEncoder().encodeToString(encryptedBytes);
  }
  
  public static String decode(String input, String salt) throws Exception {
    byte[] inputBytes = Base64.getDecoder().decode(input);
    byte[] saltBytes = Base64.getDecoder().decode(salt);
    byte[] keyBytes = new byte[16];
    byte[] ivBytes = new byte[16];
    System.arraycopy(saltBytes, 0, keyBytes, 0, 16);
    System.arraycopy(saltBytes, 16, ivBytes, 0, 16);
    SecretKeySpec keySpec = new SecretKeySpec(keyBytes, "AES");
    IvParameterSpec ivSpec = new IvParameterSpec(ivBytes);
    Cipher cipher = Cipher.getInstance("AES/CBC/PKCS5Padding");
    cipher.init(Cipher.DECRYPT_MODE, keySpec, ivSpec);
    byte[] decryptedBytes = cipher.doFinal(inputBytes);
    return new String(decryptedBytes, "UTF-8");
  }
}
```

