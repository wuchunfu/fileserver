<template>
  <div class="home-container">
    <el-menu
      :default-active="activeIndex"
      style="width: 100%; height: 50px; padding-left: 200px;"
      mode="horizontal"
      @select="handleSelect"
    >
      <el-menu-item index="1">首页</el-menu-item>
      <el-menu-item index="2">
        <el-dropdown :show-timeout="70" :hide-timeout="50" trigger="click" @command="onLanguageChange">
          <span class="el-dropdown-link">
          {{ $t("language") }}<i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="zh-cn" :disabled="disabledI18n === 'zh-cn'">简体中文</el-dropdown-item>
              <el-dropdown-item command="en" :disabled="disabledI18n === 'en'">English</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </el-menu-item>
    </el-menu>

    <el-container>
      <div id="app" style="margin: 24px 200px 24px 200px;">
        <el-row :gutter="24">
          <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24">
            <div style="display: block; margin: 24px auto; font-size: 24px; text-align: center;">
              文件共享系统
            </div>
            <div style="float: left; padding-bottom: 10px;">
              <el-button type="primary" :disabled="folderDisabled" @click="handleChangeFolder('all')">
                <icon-home theme="filled" size="10" fill="#FFFFFF"/>
                主页
              </el-button>
              <el-button type="primary" :disabled="folderDisabled" @click="handleChangeFolder('previous')">
                <icon-back theme="filled" size="10" fill="#FFFFFF"/>
                上一级
              </el-button>
              <el-button type="primary" @click="folderDialogVisible = true">
                <i class="el-icon-folder-add el-icon--right"></i> 新建文件夹
              </el-button>
              <el-button type="primary" @click="dialogVisible = true">
                <i class="el-icon-upload el-icon--right"></i> 文件上传
              </el-button>
              <el-button type="danger" :disabled="deleteDisabled" @click="deleteTips">
                <i class="el-icon-delete"></i> 批量删除
              </el-button>
            </div>

            <el-table
              ref="multipleTable"
              border
              :data="dataList"
              tooltip-effect="dark"
              style="width: 100%; margin-top: 10px;"
              @selection-change="handleSelectionChange"
              @cell-click="handleListFolder"
            >
              <el-table-column align="center" type="selection" width="50px"></el-table-column>
              <el-table-column align="left" prop="fileName" label="文件名称">
                <template #default="scope">
                  <div v-if="scope.row.isFile">
                    <span v-if="scope.row.fileType === 'txt'">
                      <icon-file-txt theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'pdf'">
                      <icon-file-pdf theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'docx'">
                      <icon-file-word theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'xlsx'">
                      <icon-file-excel theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'pptx'">
                      <icon-file-ppt theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'image'">
                      <icon-image-files theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'video'">
                      <icon-video-file theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'audio'">
                      <icon-file-music theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                    <span v-if="scope.row.fileType === 'zip'">
                      <icon-file-zip theme="filled" size="20" fill="#409EFF"/>
                      <span style="margin-left: 5px">{{ scope.row.fileName }}</span>
                    </span>
                  </div>
                  <div v-else>
                    <icon-folder-close theme="filled" size="20" fill="#409EFF"/>
                    <span style="margin-left: 5px;">{{ scope.row.fileName }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column align="center" prop="fileSize" label="文件大小"></el-table-column>
              <el-table-column align="center" prop="dateTime" label="修改时间"></el-table-column>
              <el-table-column align="center" width="100px" label="操作">
                <template v-slot="scope">
                  <el-button
                    size="mini"
                    type="primary"
                    icon="el-icon-download"
                    @click="handleDownload(scope.$index, scope.row)"
                  >
                    下载
                  </el-button>
                </template>
              </el-table-column>
            </el-table>

            <el-dialog
              v-model="folderDialogVisible"
              title="文件夹创建"
              width="30%"
              center
              :before-close="handleClose"
              :destroy-on-close="true"
            >
              <el-input v-model="folderName" @input="createFolderChange" placeholder="请输入文件夹名称" clearable/>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="handleCancelDialog">取 消</el-button>
                  <el-button type="primary" :disabled="createFolderDisabled" @click="handleCreateFolder">
                    <i class="el-icon-folder-add el-icon--right"></i> 新建
                  </el-button>
                </span>
              </template>
            </el-dialog>

            <el-dialog
              v-model="dialogVisible"
              title="文件上传"
              width="30%"
              center
              :before-close="handleClose"
              :destroy-on-close="true"
            >
              <el-upload
                ref="upload"
                drag
                action
                :http-request="handleUpload"
                multiple
                :on-change="handleChange"
                :before-remove="beforeRemove"
                :on-remove="handleRemove"
                :auto-upload="false"
                :show-file-list="true"
                :file-list="fileList"
              >
                <i class="el-icon-upload"></i>
                <div class="el-upload__text">
                  将文件拖到此处，或<em>点击上传</em>
                </div>
              </el-upload>
              <el-progress v-if="uploadFlag === true" :percentage="uploadPercent"></el-progress>
              <template #footer>
                <span class="dialog-footer">
                  <el-button @click="handleCancelDialog">取 消</el-button>
                  <el-button type="primary" :disabled="uploadDisabled" @click="uploadSubmit">
                    <i class="el-icon-upload el-icon--right"></i> 上传
                  </el-button>
                </span>
              </template>
            </el-dialog>
          </el-col>
        </el-row>
      </div>
    </el-container>
  </div>
</template>

<script lang="ts">

import { defineComponent, onMounted, reactive, ref, toRefs } from "vue";
import { i18n } from '/@/i18n';
import { ElMessage, ElMessageBox } from "element-plus";
import { AxiosRequestConfig } from "axios";
import { deleteData, getData, postData, uploadData } from "/@/api";
import { Local } from "/@/utils/storage";

export default defineComponent({
  name: 'home',
  setup() {
    const state = reactive({
      activeIndex: '1',
      disabledI18n: 'zh-cn',
      folderDialogVisible: false,
      folderName: "",
      createFolderDisabled: true,
      dialogVisible: false,
      fileParentPath: "",
      folderDisabled: true,
      dataList: [] as any,
      multipleSelection: [] as Array<string>,
      fileList: [] as Array<string>,
      uploadFlag: false,
      uploadPercent: 0,
      deleteDisabled: true,
      uploadDisabled: true,
      upload: ref(null) as any,
    });

    const handleSelect = (key: string, keyPath: Array<string>) => {
      console.log(key, keyPath)
    }

    onMounted(() => {
      const lang = Local.get("language");
      i18n.global.locale = lang;
      state.disabledI18n = lang != null ? lang : state.disabledI18n;
      getTableData();
    });

    // 语言切换
    const onLanguageChange = (lang: string) => {
      Local.set("language", lang)
      i18n.global.locale = lang;
      state.disabledI18n = lang;
    };

    const handleChange = (file: any, fileList: Array<string>) => {
      state.uploadDisabled = fileList.length === 0;
      state.fileList = fileList;
    };
    const beforeRemove = (file: any, fileList: Array<string>) => {
      return ElMessageBox.confirm(`确定移除 ${ file.name }？`);
    };
    const handleRemove = (file: any, fileList: Array<string>) => {
      state.uploadDisabled = fileList.length === 0;
    };
    const handleSelectionChange = (selection: Array<string>) => {
      state.deleteDisabled = selection.length === 0;
      state.multipleSelection = selection.map((item: any) => JSON.parse(JSON.stringify(item)).fileName);
    };

    const getTableData = () => {
      const params = {
        parentPath: state.fileParentPath
      }
      getData('/list', params).then((res: any) => {
        // console.log(res.data);
        state.dataList = res.data.data;
      }).catch((res: any) => {
        console.log(res);
        ElMessage({
          type: 'error',
          message: '获取数据失败',
        });
      });
    };

    const uploadSubmit = () => {
      state.upload?.submit();
      state.upload?.clearFiles();
    };

    const handleUpload = (fileList: any) => {
      let formData = new FormData();
      formData.set("formDataFile", fileList.file);
      formData.set("storageAbsPath", state.fileParentPath);
      let config: AxiosRequestConfig = {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
        onUploadProgress: (event: any) => {
          if (event.lengthComputable) {
            state.uploadFlag = true;
            state.uploadPercent = Number((event.loaded / event.total * 100).toFixed(0));
          }
        },
      }

      uploadData('/upload', formData, config).then((res: any) => {
        // console.log(res);
        if (state.uploadPercent === 100) {
          setTimeout(() => {
            state.uploadFlag = false;
            state.uploadPercent = 0;
            state.dialogVisible = false;
            state.uploadDisabled = true;
            state.fileList = [];
          }, 1000)
        }
        getTableData();
      }).catch((res: any) => {
        console.log(res);
        ElMessage({
          type: 'error',
          message: '上传失败!'
        });
      });
    };

    const handleDownload = (index: number, rows: any) => {
      const fileName = rows.fileName;
      let config: AxiosRequestConfig = {
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json; charset=utf-8',
          'withCredentials': 'true'
        },
        // 表明返回服务器返回的数据类型
        // 表示接收的数据为二进制文件流
        responseType: 'blob',
      }

      getData(`/download/${ fileName }`, config).then((res: any) => {
        // console.log(res);
        const blob = new Blob([res.data], { type: 'application/octet-stream' })
        // 非IE下载
        // 在拿到数据流之后,把流转为指定文件格式并创建a标签,模拟点击下载,实现文件下载功能
        // 通过 FileReader 接受并解析, 读取文件
        let reader = new FileReader();
        // 把读取的Blob和File对象以data：URL的形式返回，它与readAsArrayBuffer方法相同
        reader.readAsDataURL(blob);
        // 加载监听
        reader.onloadend = (e: any) => {
          let link = document.createElement('a');
          link.style.display = 'none';
          link.href = e.target.result;
          link.setAttribute("download", decodeURI(fileName));
          // 兼容：某些浏览器不支持HTML5的download属性
          if (typeof link.download === 'undefined') {
            link.setAttribute('target', '_blank');
          }
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);
        }
      }).catch((res: any) => {
        console.log(res);
        ElMessage({
          type: 'error',
          message: '下载失败!'
        });
      });
    };

    const handleBatchDelete = () => {
      const fileNameList = state.multipleSelection;
      const params = {
        fileNameList: fileNameList
      }

      deleteData('/delete', params).then((res: any) => {
        // console.log(res.data);
        ElMessage({
          type: 'success',
          message: '删除成功!'
        });
        getTableData();
      }).catch((res: any) => {
        console.log(res);
        ElMessage({
          type: 'error',
          message: '删除失败!'
        });
      });
    };

    const createFolderChange = () => {
      state.createFolderDisabled = state.folderName.length <= 0
    }

    const handleCreateFolder = () => {
      const params = {
        storageAbsPath: state.fileParentPath,
        folderName: state.folderName
      }
      postData('/createFolder', params).then((res: any) => {
        // console.log(res.data);
        if (res !== undefined && res.data.code === 200) {
          setTimeout(() => {
            state.folderDialogVisible = false;
            state.createFolderDisabled = true;
            state.folderName = "";
          }, 1000)
        } else {
          ElMessage({
            type: 'error',
            message: '文件夹已存在',
          });
        }
        getTableData()
      }).catch((res: any) => {
        console.log(res);
        ElMessage({
          type: 'error',
          message: '文件夹创建失败',
        });
      });
    }

    const handleListFolder = (row: any, column: any, cell: any, event: any) => {
      const data = JSON.parse(JSON.stringify(row));
      // console.log(data)
      const isFile = data.isFile;
      if (!isFile) {
        state.folderDisabled = false
        state.fileParentPath = data.filePath
        getTableData()
      }
    };

    const handleChangeFolder = (flag: string) => {
      if ("all" === flag) {
        state.fileParentPath = "";
        state.folderDisabled = true;
        getTableData()
      } else if ("previous" === flag) {
        const params = {
          parentPath: state.fileParentPath
        }
        getData('/changeFolder', params).then((res: any) => {
          // console.log(res.data);
          let data = JSON.parse(JSON.stringify(res.data.data[0]));
          if (data.parentPath === data.basePath) {
            state.folderDisabled = true
          }
          state.fileParentPath = data.parentPath
          state.dataList = res.data.data;
        }).catch((res: any) => {
          console.log(res);
          ElMessage({
            type: 'error',
            message: '文件夹切换失败',
          });
        });
      }
    };

    const handleCancelDialog = () => {
      state.folderDialogVisible = false;
      state.createFolderDisabled = true;
      state.dialogVisible = false;
      state.uploadDisabled = true;
      state.fileList = [];
      state.folderName = "";
    }

    // 提示信息
    const deleteTips = () => {
      ElMessageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        handleBatchDelete();
      }).catch(() => {
        state.multipleSelection = [];
        getTableData();
        ElMessage({
          type: 'info',
          message: '已取消删除'
        });
      });
    };

    const handleClose = (done: any) => {
      ElMessageBox.confirm('确认关闭？').then(() => {
        state.fileList = [];
        state.uploadDisabled = true;
        state.folderName = "";
        state.createFolderDisabled = true;
        done();
        ElMessage({
          type: 'success',
          message: '关闭成功!'
        });
      }).catch(() => {
        ElMessage({
          type: 'info',
          message: '已取关闭'
        });
      });
    };

    return {
      ...toRefs(state),
      handleSelect,
      onLanguageChange,
      handleChange,
      beforeRemove,
      handleRemove,
      handleSelectionChange,
      getTableData,
      uploadSubmit,
      handleUpload,
      handleDownload,
      handleBatchDelete,
      createFolderChange,
      handleCreateFolder,
      handleListFolder,
      handleChangeFolder,
      handleCancelDialog,
      deleteTips,
      handleClose
    };
  },
});
</script>

<style lang="scss" scoped>
.home-container {

}

</style>
