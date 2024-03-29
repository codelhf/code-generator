export default {
  route: {
    dashboard: '首页',
    project: '项目管理',
    projectGenerate: '生成代码',
    template: '模板管理',
    templateDetail: '模板详情',
    templateField: '模板字段',
    typeColumn: '字段映射'
  },
  navbar: {
    github: '项目地址',
    logOut: '退出登录'
  },
  tagsView: {
    refresh: '刷新',
    close: '关闭',
    closeOthers: '关闭其它',
    closeAll: '关闭所有'
  },
  common: {
    message: {
      exists: '已存在'
    },
    confirm: {
      title: '提示',
      deleteOne: '确认要删除吗？',
      cancel: '取消',
      confirm: '确定'
    },
    form: {
      cancel: '取消',
      confirm: '确定'
    }
  },
  components: {
    langSelect: {
      success: '切换语言成功'
    },
    treeList: {
      addItem: '新增',
      updateItem: '修改',
      deleteItem: '删除'
    },
    uploadExcel: {
      drag: '将文件拖到此处，或',
      browse: '点击上传',
      error: '上传头像图片大小不能超过 10MB!',
      accept: '只能上传xlsx/xls文件',
      importExcel: '导入Excel',
      formCancel: '取消',
      formConfirm: '确定'
    }
  },
  login: {
    title: '系统登录',
    logIn: '登录',
    username: '账号',
    password: '密码',
    any: '随便填',
    validUsername: '账号不能少于5位数',
    validPassword: '密码不能少于6位数'
  },
  home: {
  },
  project: {
    listQuery: {
      name: '项目名称',
      placeholderName: '请输入项目名称',
      desc: '项目描述',
      placeholderDesc: '请输入项目描述'
    },
    listButton: {
      search: '搜索',
      reset: '重置',
      add: '新增项目'
    },
    table: {
      name: '项目名称',
      desc: '项目描述',
      cteTime: '创建时间',
      uteTime: '更新时间',
      operation: '操作'
    },
    item: {
      addTitle: '新增项目',
      editTitle: '编辑项目',
      name: '项目名称',
      placeholderName: '请输入项目名称',
      desc: '项目描述',
      placeholderDesc: '请输入项目描述',
      author: '代码作者',
      placeholderAuthor: '请输入代码作者',
      organization: '所属组织',
      placeholderOrganization: '请输入所属组织: com.xxx',
      dateFormat: '日期格式',
      placeholderDateFormat: '请输入日期格式',
      fileEncode: '文件编码',
      placeholderFileEncode: '请输入文件编码',
      httpPrefix: '接口前缀',
      placeholderHttpPrefix: '请输入接口前缀',
      generateRemark: '代码注释',
      placeholderGenerateRemark: '请选择是否生成代码注释',
      generateRemark1: '是',
      generateRemark2: '否'
    },
    itemRules: {
      name: '项目名称不能为空',
      desc: '项目描述不能为空',
      author: '代码作者不能为空',
      organization: '所属组织不能为空',
      dateFormat: '日期格式不能空',
      fileEncode: '文件编码不能为空',
      httpPrefix: '接口前缀不能为空',
      generateRemark: '请选择是否生成代码注释'
    },
    database: {
      addTitle: '新增数据库',
      editTitle: '编辑数据库',
      type: '数据库类型',
      placeholderType: '请选择数据库类型',
      host: '主机',
      placeholderHost: '请输入数据库主机',
      port: '端口',
      placeholderPort: '请输入数据库主机端口',
      username: '用户',
      placeholderUsername: '请输入数据库的用户',
      password: '密码',
      placeholderPassword: '请输入数据库的密码',
      database: '数据库名',
      placeholderDatabase: '请输入数据库名称',
      delimitKeyword: '关键字处理',
      placeholderDelimitKeyword: '请输入关键字处理字符',
      languageType: '语言类型',
      placeholderLanguageType: '选择后端语言类型'
    },
    databaseRules: {
      type: '请选择数据库类型',
      url: '连接数据库URL不能为空',
      database: '数据库名称不能为空',
      host: '数据库主机不能为空',
      port: '数据库主机端口不能为空',
      username: '数据库用户不能为空',
      password: '数据库密码不能为空',
      delimitKeyword: '请输入关键字处理字符',
      languageType: '后端语言类型不能为空'
    }
  },
  projectGenerate: {
    currentTable: {
      domainName: '类名',
      placeholderDomainName: '请输入类名',
      domainDesc: '类描述',
      placeholderDomainDesc: '请输入类描述'
    },
    button: {
      tableWarning: '请选择数据表',
      templateWarning: '请添加模板',
      generateCode: '生成代码',
      addTemplate: '添加模板'
    }
  },
  projectTemplate: {
    table: {
      name: '模板名称',
      directory: '目录',
      packageName: '包名',
      fileSuffix: '后缀名',
      isGenerate: '是否生成',
      isOverride: '是否重写',
      operation: '请输入关键字搜索'
    },
    switch: {
      activeText: '是',
      inactiveText: '否'
    },
    item: {
      addTitle: '新增模板配置',
      editTitle: '编辑模板配置',
      groupId: '模板分组',
      placeholderGroup: '请选择模板分组',
      templateId: '模板名称',
      placeholderName: '请选择模板名称',
      directory: '代码目录',
      placeholderDirectory: '请输入代码目录',
      packageName: '代码包名',
      placeholderPackageName: '请输入代码包名 com.xxx',
      fileSuffix: '类名后缀',
      placeholderFileSuffix: '请输入文件后缀名 xxController 或 xx/目录/文件名',
      isGenerate: '是否生成',
      placeholderIsGenerate: '请选择是否生成代码',
      isOverride: '是否重写',
      placeholderIsOverride: '请选择是否重写代码'
    },
    itemRules: {
      projectId: '项目不能为空',
      groupId: '模板分组不能为空',
      templateId: '模板名称不能为空',
      directory: '代码目录不能为空',
      packageName: '代码包名不能为空',
      fileSuffix: '类名后缀不能为空',
      isGenerate: '请选择是否生成代码',
      isOverride: '请选择是否重写代码'
    }
  },
  template: {
    listQuery: {
      name: '模板名称',
      placeholderName: '请输入模板名称',
      desc: '模板描述',
      placeholderDesc: '请输入模板描述'
    },
    listButton: {
      search: '搜索',
      reset: '重置',
      add: '新增模板',
      export: '导出',
      import: '导入'
    },
    excel: {
      fileName: '模板列表',
      sheetName: '模板列表',
      warning: '请选择要导出的模板'
    },
    table: {
      name: '模板名称',
      desc: '模板描述',
      type: '模板类型',
      type1: '单表模板',
      type2: '通用模板',
      type3: '模板方法',
      fileType: '文件类型',
      template: '模板内容',
      operation: '操作'
    },
    item: {
      addTitle: '新增模板',
      editTitle: '编辑模板',
      group: '模板分组',
      placeholderGroup: '请选择模板分组',
      type: '模板类型',
      type1: '单表模板',
      type2: '通用模板',
      type3: '模板方法',
      placeholderType: '请选择模板类型',
      name: '模板名称',
      placeholderName: '请输入模板英文名称',
      desc: '模板描述',
      placeholderDesc: '请输入模板描述',
      fileType: '文件类型',
      placeholderFileType: '请输入生成文件的类型',
      template: '模板内容',
      placeholderTemplate: '请输入模板内容'
    },
    itemRules: {
      group: '请选择模板分组',
      type: '请选择模板类型',
      name: '模板名称为英文且不能为空',
      desc: '模板描述不能为空',
      fileType: '请输入生成文件的类型',
      template: '模板内容不能为空'
    },
    group: {
      addTitle: '新增模板分组',
      editTitle: '编辑模板分组',
      name: '分组名称',
      placeholderName: '请输入分组名称',
      desc: '分组描述',
      placeholderDesc: '请输入分组描述'
    },
    groupRules: {
      name: '模板分组名称不能为空'
    }
  },
  templateField: {
    listQuery: {
      name: '字段名称',
      placeholderName: '请输入字段名称',
      desc: '字段描述',
      placeholderDesc: '请输入字段描述'
    },
    listButton: {
      search: '搜索',
      reset: '重置',
      add: '新增模板字段'
    },
    table: {
      name: '字段名称',
      desc: '字段描述',
      type: '字段类型',
      type1: '默认字段',
      type2: '自定义字段',
      value: '字段值',
      operation: '操作'
    },
    item: {
      addTitle: '新增模板字段',
      editTitle: '编辑模板字段',
      name: '字段名称',
      placeholderName: '请输入字段名称',
      desc: '字段描述',
      placeholderDesc: '请输入字段描述',
      value: '字段值',
      placeholderValue: '请输入字段值'
    },
    itemRules: {
      name: '字段名称不能为空',
      desc: '字段描述不能为空',
      value: '字段值不能为空'
    }
  },
  typeColumn: {
    listQuery: {
      dbType: '数据库类型',
      placeholderDbType: '请选择数据库类型',
      columnType: '数据列类型',
      placeholderColumnType: '请输入数据列类型',
      fieldType: '字段类型',
      placeholderFieldType: '请输入字段类型'
    },
    listButton: {
      search: '搜索',
      reset: '重置',
      add: '新增列类型'
    },
    table: {
      dbType: '数据库类型',
      regexpType: '匹配类型',
      columnType: '列类型',
      fieldType: '字段类型',
      jdbcType: '数据类型',
      operation: '操作'
    },
    item: {
      addTitle: '新增列类型',
      editTitle: '编辑列类型',
      dbType: '数据库类型',
      placeholderDbType: '请选择数据库类型',
      regexpType: '匹配类型',
      placeholderRegexpType: '请选择匹配类型',
      regexpType1: '正则匹配',
      regexpType2: '精确匹配',
      columnType: '数据列类型',
      placeholderColumnType: '请输入数据列类型',
      languageType: '语言类型',
      placeholderLanguageType: '选择语言类型',
      fieldType: '字段类型',
      placeholderFieldType: '请输入字段类型',
      jdbcType: '数据类型',
      placeholderJdbcType: '请输入数据类型'
    },
    itemRules: {
      dbType: '请选择数据库类型',
      regexpType: '请选择匹配类型',
      columnType: '数据列类型不能为空',
      languageType: '语言类型不能为空',
      fieldType: '字段类型不能为空',
      jdbcType: '数据类型不能为空'
    },
    language: {
      addTitle: '新增语言',
      editTitle: '编辑语言',
      name: '语言名称',
      placeholderName: '请输入语言名称',
      desc: '语言描述',
      placeholderDesc: '请输入语言描述'
    },
    languageRules: {
      name: '语言名称不能为空'
    }
  }
}
