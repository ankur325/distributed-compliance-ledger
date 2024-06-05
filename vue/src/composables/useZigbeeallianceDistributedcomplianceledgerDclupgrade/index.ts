/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/vue-query";
import { useClient } from '../useClient';
import type { Ref } from 'vue'

export default function useZigbeeallianceDistributedcomplianceledgerDclupgrade() {
  const client = useClient();
  const QueryProposedUpgrade = (name: string,  options: any) => {
    const key = { type: 'QueryProposedUpgrade',  name };    
    return useQuery([key], () => {
      const { name } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryProposedUpgrade(name).then( res => res.data );
    }, options);
  }
  
  const QueryProposedUpgradeAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryProposedUpgradeAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryProposedUpgradeAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryApprovedUpgrade = (name: string,  options: any) => {
    const key = { type: 'QueryApprovedUpgrade',  name };    
    return useQuery([key], () => {
      const { name } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryApprovedUpgrade(name).then( res => res.data );
    }, options);
  }
  
  const QueryApprovedUpgradeAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryApprovedUpgradeAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryApprovedUpgradeAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRejectedUpgrade = (name: string,  options: any) => {
    const key = { type: 'QueryRejectedUpgrade',  name };    
    return useQuery([key], () => {
      const { name } = key
      return  client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryRejectedUpgrade(name).then( res => res.data );
    }, options);
  }
  
  const QueryRejectedUpgradeAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRejectedUpgradeAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.ZigbeeallianceDistributedcomplianceledgerDclupgrade.query.queryRejectedUpgradeAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryProposedUpgrade,QueryProposedUpgradeAll,QueryApprovedUpgrade,QueryApprovedUpgradeAll,QueryRejectedUpgrade,QueryRejectedUpgradeAll,
  }
}