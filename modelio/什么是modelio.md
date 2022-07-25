Modelio是基于企业体系结构，软件开发和系统工程等主要标准的，提供广泛功能的建模解决方案。

Modelio主要用于开发人员、分析人员、软件和系统架构师，它首先是一个建模环境，支持广泛的模型和图，并提供许多服务来促进架构的建模，例如模型一致性检查功能。Modelio集成了对建模和业务流程(BPMN)的支持。这些BPMN模型可以链接到工具支持的其他标准，例如UML，以便将这些业务流程集成到更大的环境中，例如企业架构。



Module 

模块通常会定义并包含:

* UML扩展，如原型、标记类型和注释类型;
* 通过Modelio MDA Java API创建、转换和导航模型;
* 上下文属性页;
* 资源管理器中的菜单命令和工具栏和属性页中的按钮命令;
* 调色板工具创建图表元素;
* 创建、修改或删除模型元素时触发的事件处理程序;
* 参数，以便在模块行为中添加一些灵活性。













概念：

Model 

Diagram 

Command 



什么是model？跟 diagram 关系



The Model API provides Java classes to represent the model element (metaclasses instances) and methods to access the attributes and relations of these elements.



* model element  ---  metaclasses instances 一一对应

* model element 组成 diagram 

* model element 是在配置文件 profile 里进行配置的

* 每个metaclass都对应一个`java interface`。For each metaclass of the `Modelio metamodel`, a corresponding Java interface is defined
* The metamodel inheritance graph between metaclasses is cloned at the Model level.元类之间的元模型继承图是在模型级别上克隆的

