### 词法分析
###### 计算机科学中将字符序列转换为记号（token）序列的过程
定义一个结构体(词法标记)：
```
  type TokenType string
  type Token struct {
    Type TokenType // token 类型
    Literal string  // token 属性值
  }
```
设想代码需要的Tokens：
例如：

### 词法分析器
语法分析器读取输入字符流、从中识别出词法标记、最后生成不同类型的记号
例如：
sum=3+2;
（
sum	标识符
=	赋值操作符
3	数字
+	加法操作符
2	数字）


