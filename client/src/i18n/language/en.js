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
  components: {
    langSelect: {
      success: 'Switch Language Success'
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
      add: 'Add Project'
    },
    table: {
      name: 'projectName',
      desc: 'ProjectDesc',
      cteTime: 'CreateTime',
      uteTime: 'UpdateTime',
      operation: 'Operate',
      run: 'GenerateCode',
      edit: 'Edit',
      database: 'Database',
      delete: 'Delete'
    },
    confirm: {
      title: 'Prompt',
      deleteOne: 'Are you sure you want to delete？',
      cancel: 'Cancel',
      confirm: 'Confirm'
    },
    item: {
      addTitle: 'Add Project',
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
      placeholderGenerateRemark: 'Please choose whether to generate code remark',
      generateRemark1: 'Yes',
      generateRemark2: 'No',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
    },
    itemRules: {
      name: 'Project name cannot be empty',
      desc: 'Item description cannot be empty',
      author: 'Code author cannot be empty',
      organization: 'Organization cannot be empty',
      dateFormat: 'Date format cannot be empty',
      fileEncode: 'File encoding cannot be empty',
      httpPrefix: 'Api prefix cannot be empty',
      generateRemark: 'Please choose whether to generate code comments'
    },
    database: {
      addTitle: 'Add Database',
      editTitle: 'Edit Database',
      type: 'Type',
      type1: 'MySQL',
      type2: 'Oracle',
      type3: 'PostgreSQL',
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
      languageType1: 'Java',
      languageType2: 'Python',
      languageType3: 'Go',
      languageType4: 'Rust',
      placeholderLanguageType: 'Select backend language type',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
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
      tableName: 'TableName',
      domainName: 'ClassName',
      placeholderDomainName: 'Please enter the class name',
      generateKey: 'PrimaryKey',
      placeholderGenerateKey: 'Please enter the primary key'
    },
    button: {
      warning: 'Please select data table',
      generateCode: 'Generate Code',
      addTemplate: 'Add Template'
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
      operation: 'Please enter keywords to search',
      edit: 'Edit',
      delete: 'Delete'
    },
    switch: {
      activeText: 'Yes',
      inactiveText: 'No'
    },
    confirm: {
      title: 'Prompt',
      deleteOne: 'Are you sure you want to delete？',
      cancel: 'Cancel',
      confirm: 'Confirm'
    },
    item: {
      addTitle: 'Add template configuration',
      editTitle: 'Edit template configuration',
      templateId: 'Template',
      placeholderName: 'Please enter the template name',
      directory: 'Directory',
      placeholderDirectory: 'Please enter the code directory',
      packageName: 'PackageName',
      placeholderPackageName: 'Please enter the code package name: com.xxx',
      fileSuffix: 'FileSuffix',
      placeholderFileSuffix: 'Please enter the file suffix xxController or xx/directory/filename',
      isGenerate: 'IsGenerate',
      placeholderIsGenerate: 'Please choose whether to generate code',
      isOverride: 'IsOverride',
      placeholderIsOverride: 'Please choose whether to rewrite the code',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
    },
    itemRules: {
      projectId: 'not operate with a project',
      templateId: 'Template cannot be empty',
      directory: 'Code directory cannot be empty',
      packageName: 'Code package name cannot be empty',
      fileSuffix: 'Class name suffix cannot be empty',
      isGenerate: 'Please choose whether to generate code',
      isOverride: 'Please choose whether to rewrite the code'
    }
  },
  template: {
    listQuery: {
      name: 'Name',
      placeholderName: 'Please enter the template name',
      desc: 'Desc',
      placeholderDesc: 'Please enter a template description'
    },
    listButton: {
      search: 'Search',
      reset: 'Reset',
      add: 'Add Template',
      export: 'export',
      import: 'import'
    },
    excel: {
      fileName: 'Template List',
      sheetName: 'Template List',
      warning: 'Please select the template to export'
    },
    table: {
      name: 'Name',
      desc: 'Desc',
      type: 'Type',
      type1: 'Table Template',
      type2: 'Common Template',
      type3: 'Method Template',
      fileType: 'FileType',
      template: 'Template',
      cteTime: 'CreateTime',
      uteTime: 'UpdateTime',
      operation: 'Operate',
      edit: 'Edit',
      delete: 'Delete'
    },
    confirm: {
      title: 'Prompt',
      deleteOne: 'Are you sure you want to delete？',
      cancel: 'Cancel',
      confirm: 'Confirm'
    },
    item: {
      addTitle: 'Add Template',
      editTitle: 'Edit Template',
      name: 'Name',
      placeholderName: 'Please enter the English name of the template',
      desc: 'Desc',
      placeholderDesc: 'Please enter a template description',
      type: 'Type',
      type1: 'Table Template',
      type2: 'Common Template',
      type3: 'Method Template',
      placeholderType: 'Please select template type',
      fileType: 'FileType',
      placeholderFileType: 'Please enter the type of generated file',
      template: 'Template',
      placeholderTemplate: 'Please enter the template content',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
    },
    itemRules: {
      name: 'The template name is in English and cannot be empty',
      exists: 'Template name already exists',
      desc: 'Template description cannot be empty',
      type: 'Please select template type',
      fileType: 'Please enter the type of generated file',
      template: 'Template content cannot be empty'
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
      add: 'Add Template Field'
    },
    table: {
      name: 'FieldName',
      desc: 'FieldDesc',
      type: 'FieldType',
      type1: 'Default field',
      type2: 'Custom field',
      value: 'FieldValue',
      cteTime: 'CreateTime',
      uteTime: 'UpdateTime',
      operation: 'Operate',
      edit: 'Edit',
      delete: 'Delete'
    },
    confirm: {
      title: 'Prompt',
      deleteOne: 'Are you sure you want to delete？',
      cancel: 'Cancel',
      confirm: 'Confirm'
    },
    item: {
      addTitle: 'Add Template Field',
      editTitle: 'Edit Template Field',
      name: 'FieldName',
      placeholderName: 'Please enter the field name',
      desc: 'FieldDesc',
      placeholderDesc: 'Please enter a field description',
      value: 'FieldValue',
      placeholderValue: 'Please enter field value',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
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
      languageType: 'Language',
      placeholderLanguageType: 'Please select language',
      fieldType: 'FieldType',
      placeholderFieldType: 'Please enter the field type'
    },
    listButton: {
      search: 'Search',
      reset: 'Reset',
      add: 'Add Field Mapping'
    },
    table: {
      dbType: 'Database',
      columnType: 'ColumnType',
      languageType: 'Language',
      fieldType: 'FieldType',
      cteTime: 'CreateTime',
      uteTime: 'UpdateTime',
      operation: 'Operate',
      edit: 'Edit',
      delete: 'Delete'
    },
    confirm: {
      title: 'Prompt',
      deleteOne: 'Are you sure you want to delete？',
      cancel: 'Cancel',
      confirm: 'Confirm'
    },
    item: {
      addTitle: 'Add Column Type',
      editTitle: 'Edit Column Type',
      dbType: 'Database',
      placeholderDbType: 'Please select database type',
      columnType: 'ColumnType',
      placeholderColumnType: 'Please enter the data column type',
      languageType: 'Language',
      placeholderLanguageType: 'Please choose language type',
      fieldType: 'FieldType',
      placeholderFieldType: 'Please enter the field type',
      formCancel: 'Cancel',
      formConfirm: 'Confirm'
    },
    itemRules: {
      dbType: 'Please select database type',
      columnType: 'Data column type cannot be empty',
      languageType: 'Language type cannot be empty',
      fieldType: 'Field type cannot be empty'
    }
  }
}
