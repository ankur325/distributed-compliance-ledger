// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgApproveAddX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRemoveNocX509IcaCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRevokeNocX509IcaCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgDeletePkiRevocationDistributionPoint } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddNocX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRemoveNocX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRejectAddX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRevokeNocX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgProposeAddX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddPkiRevocationDistributionPoint } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgProposeRevokeX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAddNocX509IcaCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgAssignVid } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRemoveX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgRevokeX509Cert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgApproveRevokeX509RootCert } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";
import { MsgUpdatePkiRevocationDistributionPoint } from "./types/zigbeealliance/distributedcomplianceledger/pki/tx";

import { ApprovedCertificates as typeApprovedCertificates} from "./types"
import { ApprovedCertificatesBySubject as typeApprovedCertificatesBySubject} from "./types"
import { ApprovedCertificatesBySubjectKeyId as typeApprovedCertificatesBySubjectKeyId} from "./types"
import { ApprovedRootCertificates as typeApprovedRootCertificates} from "./types"
import { Certificate as typeCertificate} from "./types"
import { CertificateIdentifier as typeCertificateIdentifier} from "./types"
import { ChildCertificates as typeChildCertificates} from "./types"
import { Grant as typeGrant} from "./types"
import { NocIcaCertificates as typeNocIcaCertificates} from "./types"
import { NocRootCertificates as typeNocRootCertificates} from "./types"
import { NocRootCertificatesByVidAndSkid as typeNocRootCertificatesByVidAndSkid} from "./types"
import { PkiRevocationDistributionPoint as typePkiRevocationDistributionPoint} from "./types"
import { PkiRevocationDistributionPointsByIssuerSubjectKeyID as typePkiRevocationDistributionPointsByIssuerSubjectKeyID} from "./types"
import { ProposedCertificate as typeProposedCertificate} from "./types"
import { ProposedCertificateRevocation as typeProposedCertificateRevocation} from "./types"
import { RejectedCertificate as typeRejectedCertificate} from "./types"
import { RevokedCertificates as typeRevokedCertificates} from "./types"
import { RevokedNocRootCertificates as typeRevokedNocRootCertificates} from "./types"
import { RevokedRootCertificates as typeRevokedRootCertificates} from "./types"
import { UniqueCertificate as typeUniqueCertificate} from "./types"

export { MsgApproveAddX509RootCert, MsgRemoveNocX509IcaCert, MsgRevokeNocX509IcaCert, MsgDeletePkiRevocationDistributionPoint, MsgAddNocX509RootCert, MsgRemoveNocX509RootCert, MsgAddX509Cert, MsgRejectAddX509RootCert, MsgRevokeNocX509RootCert, MsgProposeAddX509RootCert, MsgAddPkiRevocationDistributionPoint, MsgProposeRevokeX509RootCert, MsgAddNocX509IcaCert, MsgAssignVid, MsgRemoveX509Cert, MsgRevokeX509Cert, MsgApproveRevokeX509RootCert, MsgUpdatePkiRevocationDistributionPoint };

type sendMsgApproveAddX509RootCertParams = {
  value: MsgApproveAddX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgRemoveNocX509IcaCertParams = {
  value: MsgRemoveNocX509IcaCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgRevokeNocX509IcaCertParams = {
  value: MsgRevokeNocX509IcaCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgDeletePkiRevocationDistributionPointParams = {
  value: MsgDeletePkiRevocationDistributionPoint,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddNocX509RootCertParams = {
  value: MsgAddNocX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgRemoveNocX509RootCertParams = {
  value: MsgRemoveNocX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddX509CertParams = {
  value: MsgAddX509Cert,
  fee?: StdFee,
  memo?: string
};

type sendMsgRejectAddX509RootCertParams = {
  value: MsgRejectAddX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgRevokeNocX509RootCertParams = {
  value: MsgRevokeNocX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgProposeAddX509RootCertParams = {
  value: MsgProposeAddX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddPkiRevocationDistributionPointParams = {
  value: MsgAddPkiRevocationDistributionPoint,
  fee?: StdFee,
  memo?: string
};

type sendMsgProposeRevokeX509RootCertParams = {
  value: MsgProposeRevokeX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgAddNocX509IcaCertParams = {
  value: MsgAddNocX509IcaCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgAssignVidParams = {
  value: MsgAssignVid,
  fee?: StdFee,
  memo?: string
};

type sendMsgRemoveX509CertParams = {
  value: MsgRemoveX509Cert,
  fee?: StdFee,
  memo?: string
};

type sendMsgRevokeX509CertParams = {
  value: MsgRevokeX509Cert,
  fee?: StdFee,
  memo?: string
};

type sendMsgApproveRevokeX509RootCertParams = {
  value: MsgApproveRevokeX509RootCert,
  fee?: StdFee,
  memo?: string
};

type sendMsgUpdatePkiRevocationDistributionPointParams = {
  value: MsgUpdatePkiRevocationDistributionPoint,
  fee?: StdFee,
  memo?: string
};


type msgApproveAddX509RootCertParams = {
  value: MsgApproveAddX509RootCert,
};

type msgRemoveNocX509IcaCertParams = {
  value: MsgRemoveNocX509IcaCert,
};

type msgRevokeNocX509IcaCertParams = {
  value: MsgRevokeNocX509IcaCert,
};

type msgDeletePkiRevocationDistributionPointParams = {
  value: MsgDeletePkiRevocationDistributionPoint,
};

type msgAddNocX509RootCertParams = {
  value: MsgAddNocX509RootCert,
};

type msgRemoveNocX509RootCertParams = {
  value: MsgRemoveNocX509RootCert,
};

type msgAddX509CertParams = {
  value: MsgAddX509Cert,
};

type msgRejectAddX509RootCertParams = {
  value: MsgRejectAddX509RootCert,
};

type msgRevokeNocX509RootCertParams = {
  value: MsgRevokeNocX509RootCert,
};

type msgProposeAddX509RootCertParams = {
  value: MsgProposeAddX509RootCert,
};

type msgAddPkiRevocationDistributionPointParams = {
  value: MsgAddPkiRevocationDistributionPoint,
};

type msgProposeRevokeX509RootCertParams = {
  value: MsgProposeRevokeX509RootCert,
};

type msgAddNocX509IcaCertParams = {
  value: MsgAddNocX509IcaCert,
};

type msgAssignVidParams = {
  value: MsgAssignVid,
};

type msgRemoveX509CertParams = {
  value: MsgRemoveX509Cert,
};

type msgRevokeX509CertParams = {
  value: MsgRevokeX509Cert,
};

type msgApproveRevokeX509RootCertParams = {
  value: MsgApproveRevokeX509RootCert,
};

type msgUpdatePkiRevocationDistributionPointParams = {
  value: MsgUpdatePkiRevocationDistributionPoint,
};


export const registry = new Registry(msgTypes);

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	const structure: {fields: Field[]} = { fields: [] }
	for (let [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgApproveAddX509RootCert({ value, fee, memo }: sendMsgApproveAddX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveAddX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveAddX509RootCert({ value: MsgApproveAddX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveAddX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRemoveNocX509IcaCert({ value, fee, memo }: sendMsgRemoveNocX509IcaCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRemoveNocX509IcaCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRemoveNocX509IcaCert({ value: MsgRemoveNocX509IcaCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRemoveNocX509IcaCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRevokeNocX509IcaCert({ value, fee, memo }: sendMsgRevokeNocX509IcaCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRevokeNocX509IcaCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRevokeNocX509IcaCert({ value: MsgRevokeNocX509IcaCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRevokeNocX509IcaCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgDeletePkiRevocationDistributionPoint({ value, fee, memo }: sendMsgDeletePkiRevocationDistributionPointParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgDeletePkiRevocationDistributionPoint: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgDeletePkiRevocationDistributionPoint({ value: MsgDeletePkiRevocationDistributionPoint.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgDeletePkiRevocationDistributionPoint: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddNocX509RootCert({ value, fee, memo }: sendMsgAddNocX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddNocX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddNocX509RootCert({ value: MsgAddNocX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddNocX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRemoveNocX509RootCert({ value, fee, memo }: sendMsgRemoveNocX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRemoveNocX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRemoveNocX509RootCert({ value: MsgRemoveNocX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRemoveNocX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddX509Cert({ value, fee, memo }: sendMsgAddX509CertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddX509Cert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddX509Cert({ value: MsgAddX509Cert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddX509Cert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRejectAddX509RootCert({ value, fee, memo }: sendMsgRejectAddX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRejectAddX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRejectAddX509RootCert({ value: MsgRejectAddX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRejectAddX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRevokeNocX509RootCert({ value, fee, memo }: sendMsgRevokeNocX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRevokeNocX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRevokeNocX509RootCert({ value: MsgRevokeNocX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRevokeNocX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgProposeAddX509RootCert({ value, fee, memo }: sendMsgProposeAddX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgProposeAddX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgProposeAddX509RootCert({ value: MsgProposeAddX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgProposeAddX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddPkiRevocationDistributionPoint({ value, fee, memo }: sendMsgAddPkiRevocationDistributionPointParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddPkiRevocationDistributionPoint: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddPkiRevocationDistributionPoint({ value: MsgAddPkiRevocationDistributionPoint.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddPkiRevocationDistributionPoint: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgProposeRevokeX509RootCert({ value, fee, memo }: sendMsgProposeRevokeX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgProposeRevokeX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgProposeRevokeX509RootCert({ value: MsgProposeRevokeX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgProposeRevokeX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAddNocX509IcaCert({ value, fee, memo }: sendMsgAddNocX509IcaCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAddNocX509IcaCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAddNocX509IcaCert({ value: MsgAddNocX509IcaCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAddNocX509IcaCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgAssignVid({ value, fee, memo }: sendMsgAssignVidParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgAssignVid: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgAssignVid({ value: MsgAssignVid.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgAssignVid: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRemoveX509Cert({ value, fee, memo }: sendMsgRemoveX509CertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRemoveX509Cert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRemoveX509Cert({ value: MsgRemoveX509Cert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRemoveX509Cert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgRevokeX509Cert({ value, fee, memo }: sendMsgRevokeX509CertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgRevokeX509Cert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgRevokeX509Cert({ value: MsgRevokeX509Cert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgRevokeX509Cert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgApproveRevokeX509RootCert({ value, fee, memo }: sendMsgApproveRevokeX509RootCertParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgApproveRevokeX509RootCert: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgApproveRevokeX509RootCert({ value: MsgApproveRevokeX509RootCert.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgApproveRevokeX509RootCert: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgUpdatePkiRevocationDistributionPoint({ value, fee, memo }: sendMsgUpdatePkiRevocationDistributionPointParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgUpdatePkiRevocationDistributionPoint: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgUpdatePkiRevocationDistributionPoint({ value: MsgUpdatePkiRevocationDistributionPoint.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgUpdatePkiRevocationDistributionPoint: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgApproveAddX509RootCert({ value }: msgApproveAddX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgApproveAddX509RootCert", value: MsgApproveAddX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveAddX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgRemoveNocX509IcaCert({ value }: msgRemoveNocX509IcaCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRemoveNocX509IcaCert", value: MsgRemoveNocX509IcaCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRemoveNocX509IcaCert: Could not create message: ' + e.message)
			}
		},
		
		msgRevokeNocX509IcaCert({ value }: msgRevokeNocX509IcaCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeNocX509IcaCert", value: MsgRevokeNocX509IcaCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRevokeNocX509IcaCert: Could not create message: ' + e.message)
			}
		},
		
		msgDeletePkiRevocationDistributionPoint({ value }: msgDeletePkiRevocationDistributionPointParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgDeletePkiRevocationDistributionPoint", value: MsgDeletePkiRevocationDistributionPoint.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgDeletePkiRevocationDistributionPoint: Could not create message: ' + e.message)
			}
		},
		
		msgAddNocX509RootCert({ value }: msgAddNocX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAddNocX509RootCert", value: MsgAddNocX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddNocX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgRemoveNocX509RootCert({ value }: msgRemoveNocX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRemoveNocX509RootCert", value: MsgRemoveNocX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRemoveNocX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgAddX509Cert({ value }: msgAddX509CertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAddX509Cert", value: MsgAddX509Cert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddX509Cert: Could not create message: ' + e.message)
			}
		},
		
		msgRejectAddX509RootCert({ value }: msgRejectAddX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRejectAddX509RootCert", value: MsgRejectAddX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRejectAddX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgRevokeNocX509RootCert({ value }: msgRevokeNocX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeNocX509RootCert", value: MsgRevokeNocX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRevokeNocX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgProposeAddX509RootCert({ value }: msgProposeAddX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgProposeAddX509RootCert", value: MsgProposeAddX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgProposeAddX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgAddPkiRevocationDistributionPoint({ value }: msgAddPkiRevocationDistributionPointParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAddPkiRevocationDistributionPoint", value: MsgAddPkiRevocationDistributionPoint.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddPkiRevocationDistributionPoint: Could not create message: ' + e.message)
			}
		},
		
		msgProposeRevokeX509RootCert({ value }: msgProposeRevokeX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgProposeRevokeX509RootCert", value: MsgProposeRevokeX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgProposeRevokeX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgAddNocX509IcaCert({ value }: msgAddNocX509IcaCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAddNocX509IcaCert", value: MsgAddNocX509IcaCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAddNocX509IcaCert: Could not create message: ' + e.message)
			}
		},
		
		msgAssignVid({ value }: msgAssignVidParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgAssignVid", value: MsgAssignVid.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgAssignVid: Could not create message: ' + e.message)
			}
		},
		
		msgRemoveX509Cert({ value }: msgRemoveX509CertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRemoveX509Cert", value: MsgRemoveX509Cert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRemoveX509Cert: Could not create message: ' + e.message)
			}
		},
		
		msgRevokeX509Cert({ value }: msgRevokeX509CertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgRevokeX509Cert", value: MsgRevokeX509Cert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgRevokeX509Cert: Could not create message: ' + e.message)
			}
		},
		
		msgApproveRevokeX509RootCert({ value }: msgApproveRevokeX509RootCertParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgApproveRevokeX509RootCert", value: MsgApproveRevokeX509RootCert.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgApproveRevokeX509RootCert: Could not create message: ' + e.message)
			}
		},
		
		msgUpdatePkiRevocationDistributionPoint({ value }: msgUpdatePkiRevocationDistributionPointParams): EncodeObject {
			try {
				return { typeUrl: "/zigbeealliance.distributedcomplianceledger.pki.MsgUpdatePkiRevocationDistributionPoint", value: MsgUpdatePkiRevocationDistributionPoint.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgUpdatePkiRevocationDistributionPoint: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseURL: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	public structure: Record<string,unknown>;
	public registry: Array<[string, GeneratedType]> = [];

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });		
		this.updateTX(client);
		this.structure =  {
						ApprovedCertificates: getStructure(typeApprovedCertificates.fromPartial({})),
						ApprovedCertificatesBySubject: getStructure(typeApprovedCertificatesBySubject.fromPartial({})),
						ApprovedCertificatesBySubjectKeyId: getStructure(typeApprovedCertificatesBySubjectKeyId.fromPartial({})),
						ApprovedRootCertificates: getStructure(typeApprovedRootCertificates.fromPartial({})),
						Certificate: getStructure(typeCertificate.fromPartial({})),
						CertificateIdentifier: getStructure(typeCertificateIdentifier.fromPartial({})),
						ChildCertificates: getStructure(typeChildCertificates.fromPartial({})),
						Grant: getStructure(typeGrant.fromPartial({})),
						NocIcaCertificates: getStructure(typeNocIcaCertificates.fromPartial({})),
						NocRootCertificates: getStructure(typeNocRootCertificates.fromPartial({})),
						NocRootCertificatesByVidAndSkid: getStructure(typeNocRootCertificatesByVidAndSkid.fromPartial({})),
						PkiRevocationDistributionPoint: getStructure(typePkiRevocationDistributionPoint.fromPartial({})),
						PkiRevocationDistributionPointsByIssuerSubjectKeyID: getStructure(typePkiRevocationDistributionPointsByIssuerSubjectKeyID.fromPartial({})),
						ProposedCertificate: getStructure(typeProposedCertificate.fromPartial({})),
						ProposedCertificateRevocation: getStructure(typeProposedCertificateRevocation.fromPartial({})),
						RejectedCertificate: getStructure(typeRejectedCertificate.fromPartial({})),
						RevokedCertificates: getStructure(typeRevokedCertificates.fromPartial({})),
						RevokedNocRootCertificates: getStructure(typeRevokedNocRootCertificates.fromPartial({})),
						RevokedRootCertificates: getStructure(typeRevokedRootCertificates.fromPartial({})),
						UniqueCertificate: getStructure(typeUniqueCertificate.fromPartial({})),
						
		};
		client.on('signer-changed',(signer) => {			
		 this.updateTX(client);
		})
	}
	updateTX(client: IgniteClient) {
    const methods = txClient({
        signer: client.signer,
        addr: client.env.rpcURL,
        prefix: client.env.prefix ?? "cosmos",
    })
	
    this.tx = methods;
    for (let m in methods) {
        this.tx[m] = methods[m].bind(this.tx);
    }
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			ZigbeeallianceDistributedcomplianceledgerPki: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;