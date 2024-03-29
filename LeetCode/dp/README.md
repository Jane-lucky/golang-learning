# 动态规划

涉及子问题的内容，会采用到动态规划
编辑距离、背包问题

根据动态规划解题步骤（问题抽象化、建立模型、寻找约束条件、判断是否满足最优性原理、找大问题与小问题的递推关系式、填表、寻找解组成）


## 0-1背包问题
有一个背包可以装物品的总重量是W，现有N个物品，每个物品w[i],v[i]，有背包装，能装的最大价值是多少？

状态转移方程:dp[i][j]:表示前i个物品，背包重量为j是能装的中最大价值

状态转移方程：dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])

dp[i-1][j]:表示当前物品不放入背包，dp[i-1][j-w[i]]+v[i]:表示当前背包放入背包

```go
dp:=[][]int{}
```

## [72_minDistance.go 编辑距离](https://leetcode.cn/problems/edit-distance/)

问题描述：

给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符

删除一个字符

替换一个字符

分析

查找规律：

word1转变成word2所需要的最少步骤：
(1)word1[i]=word2[j] min(word1[i-1],word[j-1])

(2)word1[i]!=word[j]  
①：添加一个、删除一个： 1+min(word1[i],word2[j-1]) or 1+min(word1[i-1],word2[j])
②：替换一个：1+min(word1[i-1],word2[j-1])

状态停止条件：word1或者word2的长度为0
边界条件：一个空串和一个非空串，距离等于非空串的长度

## [HJ16_maxAgree 最大满意度](https://www.nowcoder.com/practice/f9c6f980eeec43ef85be20755ddbeaf4?tpId=37&rp=1&ru=%2Fexam%2Foj%2Fta&qru=%2Fexam%2Foj%2Fta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=&judgeStatus=&tags=593&title=&gioEnter=menu)

问题描述

输入的第 1 行，为两个正整数N，m，用一个空格隔开：

（其中 N （ N<32000 ）表示总钱数， m （m <60 ）为可购买的物品的个数。）


从第 2 行到第 m+1 行，第 j 行给出了编号为 j-1 的物品的基本数据，每行有 3 个非负整数 v p q


（其中 v 表示该物品的价格（ v<10000 ）， p 表示该物品的重要度（ 1 ~ 5 ）， q 表示该物品是主件还是附件。如果 q=0 ，表示该物品为主件，如果 q>0 ，表示该物品为附件， q 是所属主件的编号）
 最大满意度：v*w（其中v是价格，w是价值）

输出描述：

 输出一个正整数，为张强可以获得的最大的满意度。

 思路分析：
 主件和附件:买主件可买对应的附件；买附件需要判断主件是否购买

 判断条件：购买物品的总价格《N   购买物品的总个数《m
 最大满意度

状态分析：dp[i][j]表示前i个物品，钱数为j时的最大满意度

情况：主件  主件+附件 主件+附件+附件
可承重：C

物品重量w[]

所有价值v[]

dp[i][j]=max(dp[i-1][j],{四种情况产生的价值})

## [HJ61_putApple 放苹果](https://www.nowcoder.com/practice/bfd8234bb5e84be0b493656e390bdebf?tpId=37&rp=1&ru=%2Fexam%2Foj%2Fta&qru=%2Fexam%2Foj%2Fta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=&judgeStatus=&tags=593&title=&gioEnter=menu)

1. 递归的方式

当n>m时:f(m,n)=f(m,m)

当n<=m:

- 至少有一个盘子没有放苹果：f(m,n)=f(m,n-1)
- 盘子均放苹果：f(m,n)=f(m-n,n)

那么,f(m,n)=f(m,n-1)+f(m-n,n)

递归出口：

当n=1时，所有苹果都必须放在一个盘子里，所以返回１；
当没有苹果可放时，定义为１种放法；

2. 动态规划
dp[i][j]:第i个盘子放第j个苹果的方法
dp[i][j]=dp[i][j-1]+dp[i-j][j]

- 如果苹果数 < 盘子数，则表示有空盘，则忽略一个盘子，在n-1个放苹果，一直递推到n==1，有一种摆法
- 如果苹果数 >= 盘子数，可以看作没有空盘
- 则可以选择忽略一个盘子，如上边做法
- 还可以选择每个盘子放一个苹果，即苹果数剩下i-j,继续递推直到j==1

## [HJ91_plan 走方格方案](https://www.nowcoder.com/practice/e2a22f0305eb4f2f9846e7d644dba09b?tpId=37&rp=1&ru=%2Fexam%2Foj%2Fta&qru=%2Fexam%2Foj%2Fta&sourceUrl=%2Fexam%2Foj%2Fta%3FtpId%3D37&difficulty=&judgeStatus=&tags=593&title=&gioEnter=menu)

n*m的棋盘格子，从左上角达到右下角的方案(只能向下或者向右走)

dp[i][j]:表示从左上角出发，到达(i,j)的总方案路线

边界赋值：当沿棋盘边界向右向左移动时，方案仅有1种，当向内部转移时：
dp[i][j]=dp[i-1][j]+dp[i][j-1]