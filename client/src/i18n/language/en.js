export default {
  route: {
    dashboard: 'Home',
    project: 'Project',
    projectGenerate: 'GenerateCode',
    template: 'Template',
    templateDetail: 'TemplateDetail',
    templateField: 'TemplateField',
    typeColumn: 'FieldMapping'
  },
  navbar: {
    github: 'Github',
    logOut: 'Log Out'
  },
  tagsView: {
    refresh: 'Refresh',
    close: 'Close',
    closeOthers: 'Close Other',
    closeAll: 'Close All'
  },
  common: {
    message: {
      exists: 'existed'
    },
    confirm: {
      title: 'prompt',
      deleteOne: 'Are you sure you want to delete？',
      cancel: 'Cancel',
      confirm: 'Confirm'
    },
    form: {
      cancel: 'Cancel',
      confirm: 'Confirm'
    }
  },
  components: {
    langSelect: {
      success: 'Switch Language Success'
    },
    treeList: {
      addItem: 'New',
      updateItem: 'Update',
      deleteItem: 'Delete'
    },
    uploadExcel: {
      drag: 'Drop excel file here or',
      browse: 'Browse',
      error: 'The size of the uploaded avatar image cannot exceed 10MB!',
      accept: 'Only supports upload xlsx/xls files',
      importExcel: 'Import Excel',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
    }
  },
  login: {
    title: 'Login Form',
    logIn: 'Login',
    username: 'Username',
    password: 'Password',
    any: 'any',
    validUsername: 'The username can not be less than 5 digits',
    validPassword: 'The password can not be less than 6 digits'
  },
  home: {
  },
  project: {
    listQuery: {
      name: 'ProjectName',
      placeholderName: 'Please enter the project name',
      desc: 'ProjectDesc',
      placeholderDesc: 'Please enter a project description'
    },
    listButton: {
      search: 'Search',
      reset: 'Reset',
      add: 'New Project'
    },
    table: {
      name: 'ProjectName',
      desc: 'ProjectDesc',
      cteTime: 'CreateTime',
      uteTime: 'UpdateTime',
      operation: 'Operate'
    },
    item: {
      addTitle: 'New Project',
      editTitle: 'Edit Project',
      name: 'ProjectName',
      placeholderName: 'Please enter the project name',
      desc: 'ProjectDesc',
      placeholderDesc: 'Please enter a project description',
      author: 'Author',
      placeholderAuthor: 'Please enter the code author',
      organization: 'Organization',
      placeholderOrganization: 'Please enter your organization: com.xxx',
      dateFormat: 'DateFormat',
      placeholderDateFormat: 'Please enter date format',
      fileEncode: 'FileEncoding',
      placeholderFileEncode: 'Please enter the file encoding',
      httpPrefix: 'ApiPrefix',
      placeholderHttpPrefix: 'Please enter the api prefix',
      generateRemark: 'CodeRemark',
      placeholderGenerateRemark: 'Please select whether to generate code remark',
      generateRemark1: 'Yes',
      generateRemark2: 'No'
    },
    itemRules: {
      name: 'Project name cannot be empty',
      desc: 'Item description cannot be empty',
      author: 'Code author cannot be empty',
      organization: 'Organization cannot be empty',
      dateFormat: 'Date format cannot be empty',
      fileEncode: 'File encoding cannot be empty',
      httpPrefix: 'Api prefix cannot be empty',
      generateRemark: 'Please select whether to generate code comments'
    },
    database: {
      addTitle: 'New Database',
      editTitle: 'Edit Database',
      type: 'Type',
      placeholderType: 'Please select database type',
      host: 'Host',
      placeholderHost: 'Please enter the database host',
      port: 'Port',
      placeholderPort: 'Please enter the database host port',
      username: 'User',
      placeholderUsername: 'Please enter the database user',
      password: 'Password',
      placeholderPassword: 'Please enter the database password',
      database: 'Database',
      placeholderDatabase: 'Please enter the database name',
      delimitKeyword: 'Keyword',
      placeholderDelimitKeyword: 'Please enter the keyword processing characters',
      languageType: 'Language',
      placeholderLanguageType: 'Select backend language type'
    },
    databaseRules: {
      type: 'Please select database type',
      url: 'The connection database URL cannot be empty',
      database: 'The database name cannot be empty',
      host: 'The database host cannot be empty',
      port: 'The database host port cannot be empty',
      username: 'Database user cannot be empty',
      password: 'The database password cannot be empty',
      delimitKeyword: 'Please enter the keyword processing characters',
      languageType: 'Backend language type cannot be empty'
    }
  },
  projectGenerate: {
    currentTable: {
      domainName: 'ClassName',
      placeholderDomainName: 'Please enter the class name',
      domainDesc: 'ClassDesc',
      placeholderDomainDesc: 'Please enter the class desc'
    },
    button: {
      tableWarning: 'Please select data table',
      templateWarning: 'Please add a template',
      generateCode: 'Generate Code',
      addTemplate: 'New Template'
    }
  },
  projectTemplate: {
    table: {
      name: 'Template',
      directory: 'Directory',
      packageName: 'PackageName',
      fileSuffix: 'FileSuffix',
      isGenerate: 'IsGenerate',
      isOverride: 'IsOverride',
      operation: 'Please enter keywords to search'
    },
    switch: {
      activeText: 'Yes',
      inactiveText: 'No'
    },
    item: {
      addTitle: 'New template configuration',
      editTitle: 'Edit template configuration',
      groupId: 'TemplateGroup',
      placeholderGroup: 'Please select the template group',
      templateId: 'Template',
      placeholderName: 'Please select the template name',
      directory: 'Directory',
      placeholderDirectory: 'Please enter the code directory',
      packageName: 'PackageName',
      placeholderPackageName: 'Please enter the code package name: com.xxx',
      fileSuffix: 'FileSuffix',
      placeholderFileSuffix: 'Please enter the file suffix xxController or xx/directory/filename',
      isGenerate: 'IsGenerate',
      placeholderIsGenerate: 'Please select whether to generate code',
      isOverride: 'IsOverride',
      placeholderIsOverride: 'Please select whether to rewrite the code'
    },
    itemRules: {
      projectId: 'not operate with a project',
      groupId: 'Template group cannot be empty',
      templateId: 'Template cannot be empty',
      directory: 'Code directory cannot be empty',
      packageName: 'Code package name cannot be empty',
      fileSuffix: 'Class name suffix cannot be empty',
      isGenerate: 'Please select whether to generate code',
      isOverride: 'Please select whether to rewrite the code'
    }
  },
  template: {
    listQuery: {
      name: 'TemplateName',
      placeholderName: 'Please enter the template name',
      desc: 'TemplateDesc',
      placeholderDesc: 'Please enter a template description'
    },
    listButton: {
      search: 'Search',
      reset: 'Reset',
      add: 'New Template',
      import: 'Import',
      export: 'Export'
    },
    excel: {
      fileName: 'Template List',
      sheetName: 'Template List',
      warning: 'Please select the template to export'
    },
    table: {
      name: 'TemplateName',
      desc: 'TemplateDesc',
      type: 'TemplateType',
      type1: 'Table Template',
      type2: 'Common Template',
      type3: 'Method Template',
      fileType: 'FileType',
      template: 'Template',
      operation: 'Operate'
    },
    item: {
      addTitle: 'New Template',
      editTitle: 'Edit Template',
      group: 'TemplateGroup',
      placeholderGroup: 'Please select a template group',
      type: 'TemplateType',
      type1: 'Table Template',
      type2: 'Common Template',
      type3: 'Method Template',
      placeholderType: 'Please select template type',
      name: 'TemplateName',
      placeholderName: 'Please enter the English name of the template',
      desc: 'TemplateDesc',
      placeholderDesc: 'Please enter a template description',
      fileType: 'FileType',
      placeholderFileType: 'Please enter the type of generated file',
      template: 'Template',
      placeholderTemplate: 'Please enter the template content'
    },
    itemRules: {
      group: 'Please select a template group',
      type: 'Please select template type',
      name: 'The template name is in English and cannot be empty',
      desc: 'Template description cannot be empty',
      fileType: 'Please enter the type of generated file',
      template: 'Template content cannot be empty'
    },
    group: {
      addTitle: 'New template grouping',
      editTitle: 'Edit template grouping',
      name: 'GroupName',
      placeholderName: 'Please enter a group name',
      desc: 'GroupDesc',
      placeholderDesc: 'Please enter a group description'
    },
    groupRules: {
      name: 'Template group name cannot be empty'
    }
  },
  templateField: {
    listQuery: {
      name: 'FieldName',
      placeholderName: 'Please enter the field name',
      desc: 'FieldDesc',
      placeholderDesc: 'Please enter a field description'
    },
    listButton: {
      search: 'Search',
      reset: 'Reset',
      add: 'New Template Field'
    },
    table: {
      name: 'FieldName',
      desc: 'FieldDesc',
      type: 'FieldType',
      type1: 'Default field',
      type2: 'Custom field',
      value: 'FieldValue',
      operation: 'Operate'
    },
    item: {
      addTitle: 'New Template Field',
      editTitle: 'Edit Template Field',
      name: 'FieldName',
      placeholderName: 'Please enter the field name',
      desc: 'FieldDesc',
      placeholderDesc: 'Please enter a field description',
      value: 'FieldValue',
      placeholderValue: 'Please enter field value'
    },
    itemRules: {
      name: 'Field name cannot be empty',
      desc: 'Field description cannot be empty',
      value: 'Field value cannot be empty'
    }
  },
  typeColumn: {
    listQuery: {
      dbType: 'Database',
      placeholderDbType: 'Please select database type',
      columnType: 'ColumnType',
      placeholderColumnType: 'Please enter the data column type',
      fieldType: 'FieldType',
      placeholderFieldType: 'Please enter the field type'
    },
    listButton: {
      search: 'Search',
      reset: 'Reset',
      add: 'New Field Mapping'
    },
    table: {
      dbType: 'Database',
      regexpType: 'RegexpType',
      columnType: 'ColumnType',
      fieldType: 'FieldType',
      jdbcType: 'DataType',
      operation: 'Operate'
    },
    item: {
      addTitle: 'New Column Type',
      editTitle: 'Edit Column Type',
      languageType: 'Language',
      placeholderLanguageType: 'Please select language type',
      dbType: 'Database',
      placeholderDbType: 'Please select database type',
      regexpType: 'RegexpType',
      placeholderRegexpType: 'Please select regexp type',
      regexpType1: 'Regexp match',
      regexpType2: 'Exact match',
      columnType: 'ColumnType',
      placeholderColumnType: 'Please enter the data column type',
      fieldType: 'FieldType',
      placeholderFieldType: 'Please enter the field type',
      jdbcType: 'DataType',
      placeholderJdbcType: 'Please enter the data type'
    },
    itemRules: {
      languageType: 'Language type cannot be empty',
      dbType: 'Please select database type',
      regexpType: 'Please select regexp type',
      columnType: 'Data column type cannot be empty',
      fieldType: 'Field type cannot be empty',
      jdbcType: 'Data type cannot be empty'
    },
    language: {
      addTitle: 'New Language',
      editTitle: 'Edit Language',
      name: 'LanguageName',
      placeholderName: 'Please enter the language name',
      desc: 'LanguageDesc',
      placeholderDesc: 'Please enter the language desc'
    },
    languageRules: {
      name: 'Language Name cannot be empty'
    }
  }
}
