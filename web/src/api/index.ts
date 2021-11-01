import { deleteRequest, getRequestById, uploadFileRequest } from '/@/utils/requestApi';
import { AxiosRequestConfig } from "axios";

export function uploadData(url: string, data?: object, config?: AxiosRequestConfig) {
  return uploadFileRequest(url, data, config)
}

export function getData(url: string, config?: AxiosRequestConfig) {
  return getRequestById(url, config)
}

export function deleteData(url: string, data?: object) {
  return deleteRequest(url, data)
}
