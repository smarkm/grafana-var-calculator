import { getBackendSrv, getTemplateSrv } from '@grafana/runtime';
import {
  CoreApp,
  DataQueryRequest,
  DataQueryResponse,
  DataSourceApi,
  DataSourceInstanceSettings,
  createDataFrame,
  FieldType,
} from '@grafana/data';

import { MyQuery, MyDataSourceOptions, DEFAULT_QUERY, DataSourceResponse } from './types';
import { lastValueFrom } from 'rxjs';

export class DataSource extends DataSourceApi<MyQuery, MyDataSourceOptions> {
  baseUrl: string;
  url: string;

  constructor(instanceSettings: DataSourceInstanceSettings<MyDataSourceOptions>) {
    var orgiurl = instanceSettings.url
    //instanceSettings.url = "http://172.16.2.3:3000/"
    console.log(orgiurl + "::::");

    super(instanceSettings);

    this.baseUrl = orgiurl!;
    instanceSettings.url = "http://172.16.2.3:3000/";
    this.url = instanceSettings.url;

  }

  getDefaultQuery(_: CoreApp): Partial<MyQuery> {
    return DEFAULT_QUERY;
  }

  filterQuery(query: MyQuery): boolean {
    // if no query has been provided, prevent the query from being executed
    return !!query.queryText;
  }

  async query(options: DataQueryRequest<MyQuery>): Promise<DataQueryResponse> {
    const { range } = options;
    const from = range!.from.valueOf();
    const to = range!.to.valueOf();

    // Return a constant for each query.
    const data = options.targets.map((target) => {
      return createDataFrame({
        refId: target.refId,
        fields: [
          { name: 'Time', values: [from, to], type: FieldType.time },
          { name: 'Value', values: [target.constant, target.constant], type: FieldType.number },
        ],
      });
    });

    return { data };
  }

  async request(url: string, params?: string) {
    var requstUrl = `${this.baseUrl}${url}${params?.length ? `?${params}` : ''}`
    if (this) {

    }
    //requstUrl = requstUrl.replace("proxy/", "");
    console.log(requstUrl + "--------------------");
    const response = getBackendSrv().fetch<DataSourceResponse>({
      url: requstUrl,
    });
    return lastValueFrom(response);
  }

  /**
   * Checks whether we can connect to the API.
   */
  async testDatasource() {
    return {
      status: 'success',
      message: 'Success',
    };
  }

  async metricFindQuery(query: string, options: any) {
    try {
      // 解析查询表达式
      query = getTemplateSrv().replace(query);
      const result = parseQuery(query);

      // 处理返回数据
      return processResult(result);
    } catch (error) {
      console.error(error);
      return [];
    }
  }
}
// 解析查询表达式
function parseQuery(query: string) {
  try {
    // 使用 eval 解析表达式
    const result = eval(query);
    return result;
  } catch (error) {
    throw new Error(`Failed to parse query: ${error}`);
  }
}
function processResult(result: any) {
  if (Array.isArray(result)) {
    return result.map(item => ({
      text: String(item),
      value: String(item),
    }));
  } else {
    return [{ text: String(result), value: String(result) }];
  }
}