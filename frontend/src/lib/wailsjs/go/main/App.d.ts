// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {file} from '../models';
import {config} from '../models';

export function AboutEvent():Promise<void>;

export function AssociateOpen():Promise<file.File>;

export function ExportHtmlEvent():Promise<void>;

export function ExportPdfEvent():Promise<void>;

export function GetActivatedTheme():Promise<string>;

export function GetPicBed():Promise<config.PicBed>;

export function GetSaved():Promise<boolean>;

export function GetThemeList():Promise<Array<config.ExtTheme>>;

export function GetWebServerPort():Promise<number>;

export function IssuesEvent():Promise<void>;

export function LoadTheme(arg1:string):Promise<config.ThemeItem>;

export function OpenFile():Promise<void>;

export function OpenFileEvent():Promise<void>;

export function OptionsEvent():Promise<void>;

export function PersonalizaEvent():Promise<void>;

export function Quit():Promise<void>;

export function QuitEvent():Promise<void>;

export function ReadLocalFile(arg1:string):Promise<file.File>;

export function SaveAsFileEvent():Promise<void>;

export function SaveFile(arg1:file.File):Promise<file.File>;

export function SaveFileEvent():Promise<void>;

export function SelectLocalFile(arg1:file.File):Promise<file.File>;

export function SetActivatedTheme(arg1:string):Promise<void>;

export function SetBackground(arg1:file.File):Promise<void>;

export function SetSaved(arg1:boolean):Promise<void>;

export function UploadPicbed(arg1:file.File):Promise<string>;

export function UpsertPicBed(arg1:config.PicBed):Promise<void>;

export function UpsertThemeItem(arg1:config.ThemeItem):Promise<void>;
