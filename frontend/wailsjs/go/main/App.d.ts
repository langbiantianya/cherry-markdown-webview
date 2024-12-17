// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {file} from '../models';

export function AssociateOpen():Promise<file.File>;

export function ExportHtmlEvent():Promise<void>;

export function ExportPdfEvent():Promise<void>;

export function GetSaved():Promise<boolean>;

export function GetWebServerPort():Promise<number>;

export function OpenFile():Promise<void>;

export function OpenFileEvent():Promise<void>;

export function Quit():Promise<void>;

export function QuitEvent():Promise<void>;

export function ReadLocalFile(arg1:string):Promise<file.File>;

export function SaveAsFileEvent():Promise<void>;

export function SaveFile(arg1:file.File):Promise<file.File>;

export function SaveFileEvent():Promise<void>;

export function SelectLocalFile(arg1:file.File):Promise<file.File>;

export function SetSaved(arg1:boolean):Promise<void>;
