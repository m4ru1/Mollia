var codePara1 = `\`\`\` go
decrypted, err := x509.DecryptPEMBlock(block, pwd)
	
if err != nil {
    return nil, fmt.Errorf("failed PEM decryption: [%s]", err)
}
\`\`\`
`

const codePara2 = `\`\`\` go
block, err := aes.NewCipher(key)
if err != nil {
    return nil, err
}
mode := cipher.NewCBCDecrypter(block, iv)
mode.CryptBlocks(encryptedKey, encryptedKey)

//un-padding
dataLen := len(encryptedKey)
padLen := int(encryptedKey[dataLen-1])

for i := 0; i < padLen; i++ {
    if int(encryptedKey[dataLen-padLen+i]) != padLen {	// index out of range
        return nil, errors.New("padding info incorrect")
    }
}
\`\`\`
`

const codePara3 = `\`\`\` go
for i := 0; i < padLen; i++ {
    if int(encryptedKey[dataLen-padLen+i]) != padLen {	// index out of range
        return nil, errors.New("padding info incorrect")
    }
}
\`\`\`
`

const codePara4 = `\`\`\` go
//un-padding
dataLen := len(encryptedKey)
padLen := int(encryptedKey[dataLen-1])
//check padLen
if dataLen <= padLen {
    return nil, errors.New("padding info incorrect")
}
for i := 0; i < padLen; i++ {
    if int(encryptedKey[dataLen-padLen+i]) != padLen {
        return nil, errors.New("padding info incorrect")
    }
}
return encryptedKey[:dataLen-padLen], nil
\`\`\`
`

const codePara5 = `\`\`\` go
//decrypt with wrong passwd
_, err = DecryptPEMBlock(block, []byte("abcd"))
if err == nil {
    t.Error("decrypt couldn't success")
    return
}
if err.Error() != "padding info incorrect" {
    t.Errorf('unexpected error, expect \"padding info incorrect\",\n actually is \"%s\"', err)
    return
}
\`\`\`
`

const codePara6 = `\`\`\` go
// If we detect a bad padding, we assume it is an invalid password.
dlen := len(data)
if dlen == 0 || dlen%ciph.blockSize != 0 {
    return nil, errors.New("x509: invalid padding")
}
last := int(data[dlen-1])
if dlen < last {	// 直接认定为错误的解密密钥
    return nil, IncorrectPasswordError
}
if last == 0 || last > ciph.blockSize {
    return nil, IncorrectPasswordError
}
\`\`\`
`

export const testData = `# 第一章 Bug是如何发现的？

## 1.1  这是一个用来测试的二级标题1
## 1.3  这是一个用来测试的二级标题1
## 1.4  这是一个用来测试的二级标题1
## 1.5  这是一个用来测试的二级标题1
## 1.6  这是一个用来测试的二级标题1

最近在做Hyperledger Fabric的国密改造工作，因为Fabric采用Golang实现，因此国密库我选用同样适用golang实现的网安国密库[ccs-gm](https://github.com/Hyperledger-TWGC/ccs-gm)。为了确保对国密证书的处理，有一部分工作是专门将Fabric源码中引用的\`crypto/x509\`替换为ccs-gm中的\`github.com/Hyperledger-TWGC/ccs-gm/x509\`。其实x509包不单单提供了对证书处理的支持，还有一系列针对密钥读取与存储的工具方法，故事由此开始。

为了检查替换带来的影响（即换完后程序会不会有些地方寄了），我跑了一遍Fabric源码中提供的测试，奇怪的地方出现了，BCCSP模块的测试基本没啥问题，唯独一个测试方法，即\`ecdsaKeyTest\`未能通过。

![测试结果界面](https://img-blog.csdnimg.cn/0bdafada106947168a39b1828082fdf0.png)

![寄了的位置](https://img-blog.csdnimg.cn/6818bc4db9dd4dc48770e5db457a7142.png)

这块的逻辑大致描述一下：在上面的测试中，该方法首先使用口令\`passwd\`对密钥的序列化数据进行加密，然后再存储。而现在，该方法尝试用错误的口令\`passw\`来反序列化并解密得到密钥内容。这个\`passw\`并不是写错了，而是该方法想证明使用错误的密钥是无法成功提取密钥数据的。

看了一下，这个方法主要是对ECDSA密钥的创建、存储与读取进行测试。密钥在读取与存储时，会进行PEM编码/解码，并且根据需求，这个编码的内容的可以是被加密过的密文。我对这个测试方法打了一下断点，可以发现程序在这个位置会报\`index out of range\`的问题：

我们进入报错的这个\`pemToPublicKey\`方法看一下，可以发现对PEM数据的解密是借助了x509包的\`DecryptPEMBlock\`实现的。

## 1.2 这是一个用来测试的二级标题2

### 1.2.1 这是一个用来测试的三级标题

${codePara1}

我们继而进入这个\`DecryptPEMBlock\`方法一探问题所在，

${codePara2}

可以看到，对PEM编码内容的加解密是借助AES密码完成的，作为分组密码，在加解密的时候明文数据如果不满足Block长度的整数倍，我们通常会对不足的部分进行填充（Padding）,填充的每一个字节的内容都相同，为填充的字节数。而解密完成后，需要剔除掉填充的内容取出原文，方法是先读取最后一个字节的数据，从而获取填充了多少字节，然后根据这一数目剔除掉尾部的填充即可，这个过程称为解填充（Un-Padding）。

而问题就出在这里！

通过打断点发现，原本填充了5个字节，那么91~95号字节的内容都应该是5。但是由于解密时使用了错误的口令，使得解密出来的最后几个字节皆是错误的随机值。比如这里，最后一个字节便是123，而整个解密后数据的总长度才只有96字节。

![因为使用了错误口令进行解密，最后一个字节的数据padLen为123，远大于填充后明文的总长度96](https://img-blog.csdnimg.cn/cd586ef2efd04e24b0c936a8923c726c.png)

**得到了错误的padLen，紧接着，该方法便利用包含该错误值的表达式获取数组元素\`encryptedKey[dataLen-padLen+i]\`，该表达式的值显然是负值，这就是\`index out of range\`的来源了。**

回顾该方法，

${codePara3}

这一部分明显对填充字节进行了校验，以确保填充部分的每一个字节都是一样的，从而可以检查出一部分解密错误导致的问题。但这显然不够，因为一旦\`padLen\` > \`dataLen\`，那么取数组元素那一块就会报错。所以应当对padLen的大小先一步进行检查。

# 二、修改代码

上面介绍了Bug的原理，如最后一段所言，修改这个Bug只需要对PadLen先一步做校验即可。

${codePara4}

除此之外，既然Fabric的测试能测出这个问题，难道CCS-GM自己的测试里没写这种情况吗？
我一瞅，确实是，他们的key测试里没有考虑使用错误密钥进行解密的情况。我就顺道加上了。

${codePara5}

# 三、对CCS-GM仓库提交PR的流程
1. 去Github上Fork一下CCS-GM的官方仓库到自己的账号上，然后将Fork过的仓库clone到本地，进行代码修改。
2. 修改完成后，使用\`git checkout -b <new-branch-name>\`创建新的分支，并将修改后的内容提交到该分支上。
3. 注意提交的时候需要给自己的提交信息的最后附上DCO签名，即\`Signed-by-name: m4ru1 <493487492@qq.com>\`类似这样子，也可以直接\`git commit -s m "Your-commit-message"\`，\`-s\` 意味着Git会自动帮你把DCO签名添加上去。
4. 将提交后的仓库push到自己的Github仓库上，在CCS-GM的官方仓库打开一个新的Pull Request，选用你更改过的分支，并描述清楚你得更改所解决的问题，然后等待仓库的核心成员将你的分支合并进他们的仓库，这样你就成为了一个contributor（指只贡献了10行代码那种哈哈哈）。
![PR提交后的页面](https://img-blog.csdnimg.cn/d282e55ef69b4f3ebc76a75d374702c6.png)

# 四、后记

后面我又想了，\`ccs-gm/x509\`有这个问题，那\`crypto/x509\`这块是咋处理的，不会也有这问题吧？我随即便查阅了\`crypto/x509\`包的文档，其中对DecryptPEMBlock的描写让我恍然大悟。

![DecryptPEMBlock的文档描述](https://img-blog.csdnimg.cn/adb128f7395142d7be7955f67fcc0c8c.png)

首先Go语言标准库已经弃用了这一方法，因为它会导致填充预言攻击。具体而言，攻击者可以修改密文的最后一个块的最后一个字节，观察解密器的反应。如果解密器返回成功解密的结果，则说明密文填充信息的长度为1，并且填充的内容为1个字节的值。如果解密器返回解密失败的错误信息，则说明填充信息的长度可能为2，且填充内容的最后一个字节的值为2。以此类推，不断尝试，攻击者可以逐步推测出填充信息的长度和内容。一旦攻击者推测出了填充信息的长度和内容，就可以开始对明文进行恢复了（还是要对密钥做暴力搜索，但是搜索次数可以通过手段进行减少）。

抛开填充预言攻击不谈，就上面那个bug，官方库在实现的时候也更为严谨，没有出现预期的问题。

${codePara6}
`;

