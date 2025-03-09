import { DataSourcePlugin } from '@grafana/data';
import { DataSource } from './datasource';
import { MyQuery, MyDataSourceOptions } from './types';

export const plugin = new DataSourcePlugin<DataSource, MyQuery, MyDataSourceOptions>(DataSource);