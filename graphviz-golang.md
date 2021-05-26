
# Graphviz

Graphviz是一个可视化工具，使用户能够通过代码方式编辑流程图。Graphviz软件编译的源文件是以DOT语言编写的dot文件。如果想要在Mac上使用Graphviz，可以通过`brew install graphviz`.

对于任何正在编辑的dot文件，可以用[GraphvizOnline](https://dreampuf.github.io/GraphvizOnline/)在线验证其流程图效果，另有网站[webgraphviz](http://www.webgraphviz.com/)提供了更丰富的示例，可以作为参考。

## DOT基础

主要的关键字有`node, edge, graph, digraph, subgraph, strict`,大小写不敏感。任何一个DOT文件都以`digraph G`开头, `digraph G{}`将被绘制为一个空图。 
不同类型的元素会有不同的属性，可以参考[Node, Edge and Graph Attributes](https://graphviz.org/doc/info/attrs.html)介绍。
* `node`为图的核心组件之一，即节点图案，`node`有非常多类型的形状，具体可以参考[这里](https://graphviz.org/doc/info/shapes.html#html)。
* `edge`为连接不同`node`的线。
* `subgraph`在Graphviz中主要有三个作用
  1. 体现图的结构，同一个`subgraph`下的`node`和`edge`元素应该组合在一起。
  2. 提供配置统一属性的上下文。例如下面的示例中`node [style=filled,color=white];`为所有的节点配置了显示风格。
  ```   
  subgraph cluster_0 {
      style=filled;
      color=lightgrey;
      node [style=filled,color=white];
      a0 -> a1 -> a2 -> a3;
      label = "process #1";
    }
  ```
  3. 被不同的画图引擎用来定义不同的布局，比如如果subgraph的名称以cluster开头，Graphviz会将subgraph记为特殊的cluster subgraph。如果布局引擎支持，引擎会将属于同一个cluster的node绘制在一起，并将cluster的整个图形包含在边界矩形内。

## 示例dot文件

本示例参考了[github issue](https://github.com/awalterschulze/gographviz/issues/59)来添加一个匿名的subgraph，使得两个原本高度不同的两个cluster能够在同一个底部对齐。

将如下内容拷贝为`/tmp/sample.dot`,执行`dot -Tjpg /tmp/sample.dot -o sample.jpg`,那么本地将得到`sample.jpg`流程图。

```
digraph G {
             center=true;
             label="this is a title";
             labelloc=t;
             newrank=true;
             "kube-system/pod1"->cluster_source_1[ color="#C0C0C0", dir=forward, minlen=1, penwidth=2.0 ];
             cluster_source_1->cluster_source_2[ color="#C0C0C0", dir=forward, minlen=1, penwidth=2.0 ];
             "192.168.226.5"->cluster_destination_2[ color="#C0C0C0", dir=back, minlen=1, penwidth=2.0 ];
             cluster_source_2->cluster_destination_2[ color="#C0C0C0", constraint=false, dir=forward, penwidth=2.0 ];
             subgraph cluster_source {
             bgcolor="#F8F8FF";
             label="k8s-0-1";
             labeljust=l;
             style="filled,bold";
             "kube-system/pod1" [ color="#808080", fillcolor="#C8C8C8", style="filled,bold" ];
             cluster_source_1 [ color="#696969", fillcolor="#DCDCDC", label="SpoofGuard
     Forwarded", shape=box, style="rounded,filled,solid" ];
             cluster_source_2 [ color="#696969", fillcolor="#DCDCDC", label="Forwarding
     Output
     Forwarded
     Tunnel Destination IP : 172.17.1.3", shape=box, style="rounded,filled,solid" ];

     }
     ;
             subgraph cluster_destination {
             bgcolor="#F8F8FF";
             label="k8s-0-2";
             labeljust=r;
             style="filled,bold";
             "192.168.226.5" [ color="#808080", fillcolor="#C8C8C8", style="filled,bold" ];
             cluster_destination_2 [ color="#696969", fillcolor="#DCDCDC", label="Forwarding
     Received", shape=box, style="rounded,filled,solid" ];

     }
     ;
             subgraph force_node_same_level {
             rank=same;
             cluster_destination_2 [ color="#696969", fillcolor="#DCDCDC", label="Forwarding
     Received", shape=box, style="rounded,filled,solid" ];
             cluster_source_2 [ color="#696969", fillcolor="#DCDCDC", label="Forwarding
     Output
     Forwarded
     Tunnel Destination IP : 172.17.1.3", shape=box, style="rounded,filled,solid" ];
     }
     ;
     }
```
## Golang Graphviz

目前从google搜索出来的基于Go语言编写的graphviz工具有两个：
- [gographviz](https://github.com/awalterschulze/gographviz)主要提供解析DOT语言的接口和方法，能够帮助创建新图或者编辑已有的图，具体的使用可以借鉴和参考[Antrea Traceflow](https://github.com/antrea-io/antrea/blob/main/pkg/graphviz/traceflow.go)部分。
- [go-graphviz](https://github.com/goccy/go-graphviz)除了提供接口来解析DOT文件，还能够渲染流程图。

## 参考

* [Graphviz + Golang](https://levelup.gitconnected.com/graphviz-golang-the-geeky-combo-for-technical-brainstorming-76d8992d5812)
* [Graphviz中文](https://graphviztutorial.readthedocs.io/zh_CN/latest/index.html)
* [使用Graphviz画流程图](https://www.cnblogs.com/CoolJie/archive/2012/07/17/graphviz.html)