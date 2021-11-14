import { deleteRequest, getRequest, postRequest, uploadFileRequest } from '/@/utils/requestApi';
import { AxiosRequestConfig } from "axios";

export function uploadData(url: string, data?: object, config?: AxiosRequestConfig) {
  return uploadFileRequest(url, data, config)
}

export function postData(url: string, data?: object, config?: AxiosRequestConfig) {
  return postRequest(url, data, config)
}

export function getData(url: string, params?: object, config?: AxiosRequestConfig) {
  return getRequest(url, params, config)
}

export function deleteData(url: string, data?: object) {
  return deleteRequest(url, data)
}
