    本项目使用了gin+gorm框架，以mvc模式模拟实现选课系统。并使用了jwt鉴权和grbac权限管理。采用redis存储token以及通过邮箱找回密码时使用的验证码。在本地8080端口测试运行。

localhost:8080/

​	/root

​		/login：

​			Get方法，传入参数为name 、password，登录成功后返回对应角色“root”及token

​		以下管理员接口均需要携带token进行身份权限认证

​		/add_student：

​			Post方法，传入参数name、password、email，创建学生。若email重复则创建失败。

​		/add_teacher：

​			Post方法，传入参数name、password、email，创建老师。若email重复则创建失败。

​		/delete_teacher：

​			Delete方法，传入参数teacher_id，删除对应老师。若对应老师不存在则删除失败。

​		/delete_student：

​			Delete方法，传入参数student_id，删除对应学生。若对应学生不存在则删除失败。

​		/add_course ：

​			Post方法，传入参数name、weekday、class、max、teacher_id。若teacher_id对应老师不存在则创建失败。

​		/update_course：

​			Put方法，传入参数course_id、name、weekday、class、max、teacher_id。若对应课程或老师不存在则删除失败。

​	/student

​		/login：

​			Get方法，传入参数name、password。登录成功后返回对应角色“student”及token

​		以下学生接口均需要携带token进行身份权限认证

​		/check_course：

​			Get方法，无参数。返回对应学生所选择的课程。

​	/choose_course：

​			Post方法，传入参数为course_id。若课程不存在或课程时间冲突或课程人数已选满则选课失败。

​	/teacher

​		/login：

​			Get方法，传入参数为name、password，登录成功后返回对应角色“teacher”及token

​		以下老师接口均需要携带token进行身份权限认证

​		/check_student：

​			Get方法，无传入参数。返回该老师所执教的所有课程的所有学生信息



通过邮箱修改老师/学生密码：

​	/teacher(student)

​		/change_password/：

​			Get方法，传入参数为email，通过邮箱找回修改密码。若email对应的学生或老师账号不存在，则请求失败。

​			Put方法，传入参数为email，newPassword，verifycode。若验证码正确则修改密码。



​	