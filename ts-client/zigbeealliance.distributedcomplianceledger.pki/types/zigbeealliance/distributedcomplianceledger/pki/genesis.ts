/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { ApprovedCertificates } from "./approved_certificates";
import { ApprovedCertificatesBySubject } from "./approved_certificates_by_subject";
import { ApprovedCertificatesBySubjectKeyId } from "./approved_certificates_by_subject_key_id";
import { ApprovedRootCertificates } from "./approved_root_certificates";
import { ChildCertificates } from "./child_certificates";
import { NocIcaCertificates } from "./noc_ica_certificates";
import { NocRootCertificates } from "./noc_root_certificates";
import { NocRootCertificatesByVidAndSkid } from "./noc_root_certificates_by_vid_and_skid";
import { PkiRevocationDistributionPoint } from "./pki_revocation_distribution_point";
import { PkiRevocationDistributionPointsByIssuerSubjectKeyID } from "./pki_revocation_distribution_points_by_issuer_subject_key_id";
import { ProposedCertificate } from "./proposed_certificate";
import { ProposedCertificateRevocation } from "./proposed_certificate_revocation";
import { RejectedCertificate } from "./rejected_certificate";
import { RevokedCertificates } from "./revoked_certificates";
import { RevokedNocRootCertificates } from "./revoked_noc_root_certificates";
import { RevokedRootCertificates } from "./revoked_root_certificates";
import { UniqueCertificate } from "./unique_certificate";

export const protobufPackage = "zigbeealliance.distributedcomplianceledger.pki";

/** GenesisState defines the pki module's genesis state. */
export interface GenesisState {
  approvedCertificatesList: ApprovedCertificates[];
  proposedCertificateList: ProposedCertificate[];
  childCertificatesList: ChildCertificates[];
  proposedCertificateRevocationList: ProposedCertificateRevocation[];
  revokedCertificatesList: RevokedCertificates[];
  uniqueCertificateList: UniqueCertificate[];
  approvedRootCertificates: ApprovedRootCertificates | undefined;
  revokedRootCertificates: RevokedRootCertificates | undefined;
  approvedCertificatesBySubjectList: ApprovedCertificatesBySubject[];
  rejectedCertificateList: RejectedCertificate[];
  PkiRevocationDistributionPointList: PkiRevocationDistributionPoint[];
  PkiRevocationDistributionPointsByIssuerSubjectKeyIDList: PkiRevocationDistributionPointsByIssuerSubjectKeyID[];
  approvedCertificatesBySubjectKeyIdList: ApprovedCertificatesBySubjectKeyId[];
  nocRootCertificatesList: NocRootCertificates[];
  nocIcaCertificatesList: NocIcaCertificates[];
  revokedNocRootCertificatesList: RevokedNocRootCertificates[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  nocRootCertificatesByVidAndSkidList: NocRootCertificatesByVidAndSkid[];
}

function createBaseGenesisState(): GenesisState {
  return {
    approvedCertificatesList: [],
    proposedCertificateList: [],
    childCertificatesList: [],
    proposedCertificateRevocationList: [],
    revokedCertificatesList: [],
    uniqueCertificateList: [],
    approvedRootCertificates: undefined,
    revokedRootCertificates: undefined,
    approvedCertificatesBySubjectList: [],
    rejectedCertificateList: [],
    PkiRevocationDistributionPointList: [],
    PkiRevocationDistributionPointsByIssuerSubjectKeyIDList: [],
    approvedCertificatesBySubjectKeyIdList: [],
    nocRootCertificatesList: [],
    nocIcaCertificatesList: [],
    revokedNocRootCertificatesList: [],
    nocRootCertificatesByVidAndSkidList: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.approvedCertificatesList) {
      ApprovedCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.proposedCertificateList) {
      ProposedCertificate.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.childCertificatesList) {
      ChildCertificates.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.proposedCertificateRevocationList) {
      ProposedCertificateRevocation.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.revokedCertificatesList) {
      RevokedCertificates.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.uniqueCertificateList) {
      UniqueCertificate.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.approvedRootCertificates !== undefined) {
      ApprovedRootCertificates.encode(message.approvedRootCertificates, writer.uint32(58).fork()).ldelim();
    }
    if (message.revokedRootCertificates !== undefined) {
      RevokedRootCertificates.encode(message.revokedRootCertificates, writer.uint32(66).fork()).ldelim();
    }
    for (const v of message.approvedCertificatesBySubjectList) {
      ApprovedCertificatesBySubject.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    for (const v of message.rejectedCertificateList) {
      RejectedCertificate.encode(v!, writer.uint32(82).fork()).ldelim();
    }
    for (const v of message.PkiRevocationDistributionPointList) {
      PkiRevocationDistributionPoint.encode(v!, writer.uint32(90).fork()).ldelim();
    }
    for (const v of message.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList) {
      PkiRevocationDistributionPointsByIssuerSubjectKeyID.encode(v!, writer.uint32(98).fork()).ldelim();
    }
    for (const v of message.approvedCertificatesBySubjectKeyIdList) {
      ApprovedCertificatesBySubjectKeyId.encode(v!, writer.uint32(106).fork()).ldelim();
    }
    for (const v of message.nocRootCertificatesList) {
      NocRootCertificates.encode(v!, writer.uint32(114).fork()).ldelim();
    }
    for (const v of message.nocIcaCertificatesList) {
      NocIcaCertificates.encode(v!, writer.uint32(122).fork()).ldelim();
    }
    for (const v of message.revokedNocRootCertificatesList) {
      RevokedNocRootCertificates.encode(v!, writer.uint32(130).fork()).ldelim();
    }
    for (const v of message.nocRootCertificatesByVidAndSkidList) {
      NocRootCertificatesByVidAndSkid.encode(v!, writer.uint32(138).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.approvedCertificatesList.push(ApprovedCertificates.decode(reader, reader.uint32()));
          break;
        case 2:
          message.proposedCertificateList.push(ProposedCertificate.decode(reader, reader.uint32()));
          break;
        case 3:
          message.childCertificatesList.push(ChildCertificates.decode(reader, reader.uint32()));
          break;
        case 4:
          message.proposedCertificateRevocationList.push(ProposedCertificateRevocation.decode(reader, reader.uint32()));
          break;
        case 5:
          message.revokedCertificatesList.push(RevokedCertificates.decode(reader, reader.uint32()));
          break;
        case 6:
          message.uniqueCertificateList.push(UniqueCertificate.decode(reader, reader.uint32()));
          break;
        case 7:
          message.approvedRootCertificates = ApprovedRootCertificates.decode(reader, reader.uint32());
          break;
        case 8:
          message.revokedRootCertificates = RevokedRootCertificates.decode(reader, reader.uint32());
          break;
        case 9:
          message.approvedCertificatesBySubjectList.push(ApprovedCertificatesBySubject.decode(reader, reader.uint32()));
          break;
        case 10:
          message.rejectedCertificateList.push(RejectedCertificate.decode(reader, reader.uint32()));
          break;
        case 11:
          message.PkiRevocationDistributionPointList.push(
            PkiRevocationDistributionPoint.decode(reader, reader.uint32()),
          );
          break;
        case 12:
          message.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList.push(
            PkiRevocationDistributionPointsByIssuerSubjectKeyID.decode(reader, reader.uint32()),
          );
          break;
        case 13:
          message.approvedCertificatesBySubjectKeyIdList.push(
            ApprovedCertificatesBySubjectKeyId.decode(reader, reader.uint32()),
          );
          break;
        case 14:
          message.nocRootCertificatesList.push(NocRootCertificates.decode(reader, reader.uint32()));
          break;
        case 15:
          message.nocIcaCertificatesList.push(NocIcaCertificates.decode(reader, reader.uint32()));
          break;
        case 16:
          message.revokedNocRootCertificatesList.push(RevokedNocRootCertificates.decode(reader, reader.uint32()));
          break;
        case 17:
          message.nocRootCertificatesByVidAndSkidList.push(
            NocRootCertificatesByVidAndSkid.decode(reader, reader.uint32()),
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      approvedCertificatesList: Array.isArray(object?.approvedCertificatesList)
        ? object.approvedCertificatesList.map((e: any) => ApprovedCertificates.fromJSON(e))
        : [],
      proposedCertificateList: Array.isArray(object?.proposedCertificateList)
        ? object.proposedCertificateList.map((e: any) => ProposedCertificate.fromJSON(e))
        : [],
      childCertificatesList: Array.isArray(object?.childCertificatesList)
        ? object.childCertificatesList.map((e: any) => ChildCertificates.fromJSON(e))
        : [],
      proposedCertificateRevocationList: Array.isArray(object?.proposedCertificateRevocationList)
        ? object.proposedCertificateRevocationList.map((e: any) => ProposedCertificateRevocation.fromJSON(e))
        : [],
      revokedCertificatesList: Array.isArray(object?.revokedCertificatesList)
        ? object.revokedCertificatesList.map((e: any) => RevokedCertificates.fromJSON(e))
        : [],
      uniqueCertificateList: Array.isArray(object?.uniqueCertificateList)
        ? object.uniqueCertificateList.map((e: any) => UniqueCertificate.fromJSON(e))
        : [],
      approvedRootCertificates: isSet(object.approvedRootCertificates)
        ? ApprovedRootCertificates.fromJSON(object.approvedRootCertificates)
        : undefined,
      revokedRootCertificates: isSet(object.revokedRootCertificates)
        ? RevokedRootCertificates.fromJSON(object.revokedRootCertificates)
        : undefined,
      approvedCertificatesBySubjectList: Array.isArray(object?.approvedCertificatesBySubjectList)
        ? object.approvedCertificatesBySubjectList.map((e: any) => ApprovedCertificatesBySubject.fromJSON(e))
        : [],
      rejectedCertificateList: Array.isArray(object?.rejectedCertificateList)
        ? object.rejectedCertificateList.map((e: any) => RejectedCertificate.fromJSON(e))
        : [],
      PkiRevocationDistributionPointList: Array.isArray(object?.PkiRevocationDistributionPointList)
        ? object.PkiRevocationDistributionPointList.map((e: any) => PkiRevocationDistributionPoint.fromJSON(e))
        : [],
      PkiRevocationDistributionPointsByIssuerSubjectKeyIDList:
        Array.isArray(object?.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList)
          ? object.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList.map((e: any) =>
            PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromJSON(e)
          )
          : [],
      approvedCertificatesBySubjectKeyIdList: Array.isArray(object?.approvedCertificatesBySubjectKeyIdList)
        ? object.approvedCertificatesBySubjectKeyIdList.map((e: any) => ApprovedCertificatesBySubjectKeyId.fromJSON(e))
        : [],
      nocRootCertificatesList: Array.isArray(object?.nocRootCertificatesList)
        ? object.nocRootCertificatesList.map((e: any) => NocRootCertificates.fromJSON(e))
        : [],
      nocIcaCertificatesList: Array.isArray(object?.nocIcaCertificatesList)
        ? object.nocIcaCertificatesList.map((e: any) => NocIcaCertificates.fromJSON(e))
        : [],
      revokedNocRootCertificatesList: Array.isArray(object?.revokedNocRootCertificatesList)
        ? object.revokedNocRootCertificatesList.map((e: any) => RevokedNocRootCertificates.fromJSON(e))
        : [],
      nocRootCertificatesByVidAndSkidList: Array.isArray(object?.nocRootCertificatesByVidAndSkidList)
        ? object.nocRootCertificatesByVidAndSkidList.map((e: any) => NocRootCertificatesByVidAndSkid.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.approvedCertificatesList) {
      obj.approvedCertificatesList = message.approvedCertificatesList.map((e) =>
        e ? ApprovedCertificates.toJSON(e) : undefined
      );
    } else {
      obj.approvedCertificatesList = [];
    }
    if (message.proposedCertificateList) {
      obj.proposedCertificateList = message.proposedCertificateList.map((e) =>
        e ? ProposedCertificate.toJSON(e) : undefined
      );
    } else {
      obj.proposedCertificateList = [];
    }
    if (message.childCertificatesList) {
      obj.childCertificatesList = message.childCertificatesList.map((e) => e ? ChildCertificates.toJSON(e) : undefined);
    } else {
      obj.childCertificatesList = [];
    }
    if (message.proposedCertificateRevocationList) {
      obj.proposedCertificateRevocationList = message.proposedCertificateRevocationList.map((e) =>
        e ? ProposedCertificateRevocation.toJSON(e) : undefined
      );
    } else {
      obj.proposedCertificateRevocationList = [];
    }
    if (message.revokedCertificatesList) {
      obj.revokedCertificatesList = message.revokedCertificatesList.map((e) =>
        e ? RevokedCertificates.toJSON(e) : undefined
      );
    } else {
      obj.revokedCertificatesList = [];
    }
    if (message.uniqueCertificateList) {
      obj.uniqueCertificateList = message.uniqueCertificateList.map((e) => e ? UniqueCertificate.toJSON(e) : undefined);
    } else {
      obj.uniqueCertificateList = [];
    }
    message.approvedRootCertificates !== undefined && (obj.approvedRootCertificates = message.approvedRootCertificates
      ? ApprovedRootCertificates.toJSON(message.approvedRootCertificates)
      : undefined);
    message.revokedRootCertificates !== undefined && (obj.revokedRootCertificates = message.revokedRootCertificates
      ? RevokedRootCertificates.toJSON(message.revokedRootCertificates)
      : undefined);
    if (message.approvedCertificatesBySubjectList) {
      obj.approvedCertificatesBySubjectList = message.approvedCertificatesBySubjectList.map((e) =>
        e ? ApprovedCertificatesBySubject.toJSON(e) : undefined
      );
    } else {
      obj.approvedCertificatesBySubjectList = [];
    }
    if (message.rejectedCertificateList) {
      obj.rejectedCertificateList = message.rejectedCertificateList.map((e) =>
        e ? RejectedCertificate.toJSON(e) : undefined
      );
    } else {
      obj.rejectedCertificateList = [];
    }
    if (message.PkiRevocationDistributionPointList) {
      obj.PkiRevocationDistributionPointList = message.PkiRevocationDistributionPointList.map((e) =>
        e ? PkiRevocationDistributionPoint.toJSON(e) : undefined
      );
    } else {
      obj.PkiRevocationDistributionPointList = [];
    }
    if (message.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList) {
      obj.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = message
        .PkiRevocationDistributionPointsByIssuerSubjectKeyIDList.map((e) =>
          e ? PkiRevocationDistributionPointsByIssuerSubjectKeyID.toJSON(e) : undefined
        );
    } else {
      obj.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList = [];
    }
    if (message.approvedCertificatesBySubjectKeyIdList) {
      obj.approvedCertificatesBySubjectKeyIdList = message.approvedCertificatesBySubjectKeyIdList.map((e) =>
        e ? ApprovedCertificatesBySubjectKeyId.toJSON(e) : undefined
      );
    } else {
      obj.approvedCertificatesBySubjectKeyIdList = [];
    }
    if (message.nocRootCertificatesList) {
      obj.nocRootCertificatesList = message.nocRootCertificatesList.map((e) =>
        e ? NocRootCertificates.toJSON(e) : undefined
      );
    } else {
      obj.nocRootCertificatesList = [];
    }
    if (message.nocIcaCertificatesList) {
      obj.nocIcaCertificatesList = message.nocIcaCertificatesList.map((e) =>
        e ? NocIcaCertificates.toJSON(e) : undefined
      );
    } else {
      obj.nocIcaCertificatesList = [];
    }
    if (message.revokedNocRootCertificatesList) {
      obj.revokedNocRootCertificatesList = message.revokedNocRootCertificatesList.map((e) =>
        e ? RevokedNocRootCertificates.toJSON(e) : undefined
      );
    } else {
      obj.revokedNocRootCertificatesList = [];
    }
    if (message.nocRootCertificatesByVidAndSkidList) {
      obj.nocRootCertificatesByVidAndSkidList = message.nocRootCertificatesByVidAndSkidList.map((e) =>
        e ? NocRootCertificatesByVidAndSkid.toJSON(e) : undefined
      );
    } else {
      obj.nocRootCertificatesByVidAndSkidList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.approvedCertificatesList = object.approvedCertificatesList?.map((e) => ApprovedCertificates.fromPartial(e))
      || [];
    message.proposedCertificateList = object.proposedCertificateList?.map((e) => ProposedCertificate.fromPartial(e))
      || [];
    message.childCertificatesList = object.childCertificatesList?.map((e) => ChildCertificates.fromPartial(e)) || [];
    message.proposedCertificateRevocationList =
      object.proposedCertificateRevocationList?.map((e) => ProposedCertificateRevocation.fromPartial(e)) || [];
    message.revokedCertificatesList = object.revokedCertificatesList?.map((e) => RevokedCertificates.fromPartial(e))
      || [];
    message.uniqueCertificateList = object.uniqueCertificateList?.map((e) => UniqueCertificate.fromPartial(e)) || [];
    message.approvedRootCertificates =
      (object.approvedRootCertificates !== undefined && object.approvedRootCertificates !== null)
        ? ApprovedRootCertificates.fromPartial(object.approvedRootCertificates)
        : undefined;
    message.revokedRootCertificates =
      (object.revokedRootCertificates !== undefined && object.revokedRootCertificates !== null)
        ? RevokedRootCertificates.fromPartial(object.revokedRootCertificates)
        : undefined;
    message.approvedCertificatesBySubjectList =
      object.approvedCertificatesBySubjectList?.map((e) => ApprovedCertificatesBySubject.fromPartial(e)) || [];
    message.rejectedCertificateList = object.rejectedCertificateList?.map((e) => RejectedCertificate.fromPartial(e))
      || [];
    message.PkiRevocationDistributionPointList =
      object.PkiRevocationDistributionPointList?.map((e) => PkiRevocationDistributionPoint.fromPartial(e)) || [];
    message.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList =
      object.PkiRevocationDistributionPointsByIssuerSubjectKeyIDList?.map((e) =>
        PkiRevocationDistributionPointsByIssuerSubjectKeyID.fromPartial(e)
      ) || [];
    message.approvedCertificatesBySubjectKeyIdList =
      object.approvedCertificatesBySubjectKeyIdList?.map((e) => ApprovedCertificatesBySubjectKeyId.fromPartial(e))
      || [];
    message.nocRootCertificatesList = object.nocRootCertificatesList?.map((e) => NocRootCertificates.fromPartial(e))
      || [];
    message.nocIcaCertificatesList = object.nocIcaCertificatesList?.map((e) => NocIcaCertificates.fromPartial(e)) || [];
    message.revokedNocRootCertificatesList =
      object.revokedNocRootCertificatesList?.map((e) => RevokedNocRootCertificates.fromPartial(e)) || [];
    message.nocRootCertificatesByVidAndSkidList =
      object.nocRootCertificatesByVidAndSkidList?.map((e) => NocRootCertificatesByVidAndSkid.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
