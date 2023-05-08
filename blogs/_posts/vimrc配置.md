---
title: vimrc配置
mathjax: false
date: 2023-05-05 20:20:28
tags:
categories:
---

## vim配置

## 参考资料

```v
"==============================================================================
" set 一般设置
"==============================================================================
set nocompatible                "不兼容vi
syntax enable                   "语法高亮使能
syntax on                       "语法高亮开启
set hlsearch                    "开启高亮搜索
set incsearch                   "输出字符串的同时进行搜索
set ignorecase                  "搜索时忽略大小写
set ruler                       "右下角的状态栏下显示光标的坐标
set showcmd                     "显示normal模式下的命令
set number                      "显示行号
set relativenumber              "相对行号
set cmdheight=1                 "命令行高度为1行"
set noshowmode                  "不显示'-- insert --' 等
set autoindent                  "自动缩进
set smartindent                 "智能缩进
set cindent                     "c语言缩进
set pastetoggle=<F4>            "系统paste
set mouse=i                     "仅在insert mode 下可使用mouse

set autoread                    "当文件在外部被改变时, vim自动载入
set showmatch                   "显示匹配的括号
set matchtime=0                 "显示匹配的括号的时间，单位：十分之一秒
set t_Co=256                    "vim配色为256色
set cursorline                  "高亮当前光标行
set cursorcolumn                "高亮当前光标列
autocmd InsertEnter * se cul
set shortmess=atI
set scrolloff=3                 "光标移动到buffer的顶部和底部时保持3行距离
set laststatus=2                "指定什么时候对最后一个窗口显示状态条, 2表示永远有
set cinoptions=g0,:0,N-s,(0     "设置C/C++语言的具体缩进方式, unkown
set tabstop=4                   "读取到文件的\t时，解释的空白长度
set softtabstop=4               "在编辑动作（按tab或backspace键）时，所输入或删除的空白长度
set shiftwidth=4                "自动缩进时的空白长度
set expandtab                   "将tab转换成space，文件中存储的将会是space
set smarttab
" set list                      "显示tab键为^I, 行为显示$
set listchars=tab:>-,trail:-    "tab显示为">---" 行尾多余的空白字符显示为"-"
set history=1000                "命令历史个数
set langmenu=en_US.UTF-8        "菜单语系
set helplang=en                 "使用中文文档cn，英文则为en
set autowrite                   "自动保存, 否则切换buffer,文件后会会提示没有保存
set magic                       "匹配正则表达式时(\m)：除了$.*^之外其他正则表达式的元字符都要加反斜杠\
set confirm                     "总是出现确认选项，如执行:qa等命令时
set nobackup                    "vim当覆盖一个文件时不保留一个备份
set noswapfile                  "不使用交换空间
set wildmenu                    "补全相关，unkown
set completeopt=longest,menu    "显示补全列表
set backspace=indent,eol,start  "indent将4个空格缩进当缩进删除，eol删除到上一行
set report=0                    "当使用:commands命令时，总是报告文件的那一行被改变过
set fillchars=vert:\ ,stl:\ ,stlnc:\  "设置多窗口时的窗口间分割线
" set iskeyword+=.                "追加.为判断一个word
set termencoding=utf-8          "终端使用的字符编码
set encoding=utf8               "vim内部buffer使用的字符编码
set fileencodings=utf8,ucs-bom,gbk,cp936,gb2312,gb18030 "依次判断文件字符编码
" set foldmethod=indent         "基于缩进的折叠
set foldmethod=syntax           "基于语法的折叠
set nofoldenable                "启动 vim 时关闭折叠代码
" set nowrap                      "禁止折行
set viminfo+=n~/.cache/vim/.viminfo

let $BASH_ENV="~/.bash_aliases" "增加外部命令的alias
let mapleader = ","             "修改<leader>键为','
vnoremap <leader>y "+y

filetype plugin on              "根据文件类型读取插件
filetype indent on              "使用缩进文件

" 标签跳转快捷键，在gnome-terminal中无效
" normal no recursive map
" insert no recursive map
nnoremap <C-c> :tabclose<cr>
nnoremap <C-x> :tabonly<cr>
nnoremap <C-k> :tabs<cr>
nnoremap <C-n> :tabnew<cr>
nnoremap <C-h> :tabprevious<cr>
nnoremap <C-l> :tabnext<cr>

" 光标闪烁，在插入模式下为‘|’，在norrmal 模式下为方块
" 若要不闪烁，则将改为后面的数字
let &t_SI = "\<Esc>[5 q"                     " 6
let &t_SR = "\<Esc>[3 q"                     " 4
let &t_EI = "\<Esc>[1 q"                     " 2
```

很神奇我的配置就是那么简单哈哈哈哈 不想有太多有的没的东西 就是说 

```v
" basic config 
set nocompatible
set bs=indent,eol,start
set smartindent  
set relativenumber
set tabstop=2
set shiftwidth=2
set nofoldenable
set noswapfile
:syntax on 
" theme you should have a try
":colorschem ron | murphy | slate | darkblud | delek  
:colorschem delek 
" hot key
:inoremap jk <ESC>
noremap H ^
noremap L $
inoremap <C-k> <Up>
inoremap <C-j> <Down>
inoremap <C-a> <Home>
inoremap <C-e> <End>
```

