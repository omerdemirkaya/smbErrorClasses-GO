if err != nil{
    errorMessage := err.Error()
    var status string
    switch {
	case strings.Contains(errorMessage, "STATUS_NOT_IMPLEMENTED"):
        status = "Geçersiz işlem: STATUS_NOT_IMPLEMENTED error 0xC0000002, ERRbadfunc 0x0001 : Invalid Function."
    case strings.Contains(errorMessage, "STATUS_INVALID_DEVICE_REQUEST"):
        status = "Geçersiz işlem: STATUS_INVALID_DEVICE_REQUEST error 0xC0000010, ERRbadfunc 0x0001: Invalid Function."
    case strings.Contains(errorMessage, "STATUS_ILLEGAL_FUNCTION"):
        status = "Geçersiz işlem: STATUS_ILLEGAL_FUNCTION error 0xC00000AF, ERRbadfunc 0x0001: Invalid Function."
    case strings.Contains(errorMessage, "STATUS_NO_SUCH_FILE"):
        status = "Dosya bulunamadı: STATUS_NO_SUCH_FILE error 0xC000000F, ERRbadfile 0x0002: File not found."
    case strings.Contains(errorMessage, "STATUS_NO_SUCH_DEVICE"):
        status = "Dosya bulunamadı: STATUS_NO_SUCH_DEVICE error 0xC000000E, ERRbadfile 0x0002: File not found."
    case strings.Contains(errorMessage, "STATUS_OBJECT_NAME_NOT_FOUND"):
        status = "Dosya bulunamadı: STATUS_OBJECT_NAME_NOT_FOUND error 0xC0000034, ERRbadfile 0x0002: File not found."
    case strings.Contains(errorMessage, "STATUS_OBJECT_PATH_INVALID"):
        status = "Geçersiz yol: STATUS_OBJECT_PATH_INVALID error 0xC0000039, ERRbadpath 0x0003: A component in the path prefix is not a directory."
    case strings.Contains(errorMessage, "STATUS_OBJECT_PATH_SYNTAX_BAD"):
        status = "Geçersiz yol: STATUS_OBJECT_PATH_SYNTAX_BAD error 0xC000003B, ERRbadpath 0x0003: A component in the path prefix is not a directory."
    case strings.Contains(errorMessage, "STATUS_DFS_EXIT_PATH_FOUND"):
        status = "Geçersiz yol: STATUS_DFS_EXIT_PATH_FOUND error 0xC000009B, ERRbadpath 0x0003: A component in the path prefix is not a directory."
    case strings.Contains(errorMessage, "STATUS_REDIRECTOR_NOT_STARTED"):
        status = "Geçersiz yol: STATUS_REDIRECTOR_NOT_STARTED error 0xC00000FB, ERRbadpath 0x0003: A component in the path prefix is not a directory."
    case strings.Contains(errorMessage, "STATUS_TOO_MANY_OPENED_FILES"):
        status = "Açılmış dosya sayısı fazla: STATUS_TOO_MANY_OPENED_FILES error 0xC000011F, ERRnofids 0x0004: Too many open files. No FIDs are available."
    case strings.Contains(errorMessage, "STATUS_ACCESS_DENIED"):
        status = "Erişim reddedildi: STATUS_ACCESS_DENIED error 0xC0000022, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_INVALID_LOCK_SEQUENCE"):
        status = "Geçersiz kilit sırası: STATUS_INVALID_LOCK_SEQUENCE error 0xC000001E, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_INVALID_VIEW_SIZE"):
        status = "Geçersiz görünüm boyutu: STATUS_INVALID_VIEW_SIZE error 0xC000001F, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_ALREADY_COMMITTED"):
        status = "Zaten taahhüt edildi: STATUS_ALREADY_COMMITTED error 0xC0000021, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_PORT_CONNECTION_REFUSED"):
        status = "Bağlantı reddedildi: STATUS_PORT_CONNECTION_REFUSED error 0xC0000041, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_THREAD_IS_TERMINATING"):
        status = "İşlem sonlanıyor: STATUS_THREAD_IS_TERMINATING error 0xC000004B, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_DELETE_PENDING"):
        status = "Silme beklemede: STATUS_DELETE_PENDING error 0xC0000056, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_PRIVILEGE_NOT_HELD"):
        status = "Yetki sahibi değil: STATUS_PRIVILEGE_NOT_HELD error 0xC0000061, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_LOGON_FAILURE"):
        status = "Giriş hatası: STATUS_LOGON_FAILURE error 0xC000006D, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_FILE_IS_A_DIRECTORY"):
        status = "Dosya bir dizin: STATUS_FILE_IS_A_DIRECTORY error 0xC00000BA, ERRnoaccess 0x0005: Access denied."
    case strings.Contains(errorMessage, "STATUS_FILE_RENAMED"):
        status = "Dosya yeniden adlandırıldı: STATUS_FILE_RENAMED error 0xC00000D5, ERRnoaccess 0x0005: Access denied."
	case strings.Contains(errorMessage, "STATUS_PROCESS_IS_TERMINATING"):
			status = "İşlem sonlanıyor: STATUS_PROCESS_IS_TERMINATING error 0xC000010A, ERRnoaccess 0x0005: Access denied."
	case strings.Contains(errorMessage, "STATUS_DIRECTORY_NOT_EMPTY"):
			status = "Klasör boş değil: STATUS_DIRECTORY_NOT_EMPTY error 0xC0000101, ERRnoaccess 0x0005: Access denied."
	case strings.Contains(errorMessage, "STATUS_SMB_BAD_FID"):
			status = "Geçersiz FID: STATUS_SMB_BAD_FID error 0xC0000121, ERRnoaccess 0x0005: Access denied."
	case strings.Contains(errorMessage, "STATUS_FILE_DELETED"):
			status = "Dosya silindi: STATUS_FILE_DELETED error 0x00060001, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_INVALID_HANDLE"):
			status = "Geçersiz handle: STATUS_INVALID_HANDLE error 0xC0000008, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_OBJECT_TYPE_MISMATCH"):
			status = "Nesne türü uyuşmazlığı: STATUS_OBJECT_TYPE_MISMATCH error 0xC0000024, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_PORT_DISCONNECTED"):
			status = "Port bağlantısı kesildi: STATUS_PORT_DISCONNECTED error 0xC0000037, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_INVALID_PORT_HANDLE"):
			status = "Geçersiz port handle: STATUS_INVALID_PORT_HANDLE error 0xC0000042, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_FILE_CLOSED"):
			status = "Dosya kapalı: STATUS_FILE_CLOSED error 0xC0000128, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_HANDLE_NOT_CLOSABLE"):
			status = "Handle kapatılamıyor: STATUS_HANDLE_NOT_CLOSABLE error 0xC0000235, ERRbadfid 0x0006: Invalid FID."
	case strings.Contains(errorMessage, "STATUS_SECTION_TOO_BIG"):
			status = "Bölüm çok büyük: STATUS_SECTION_TOO_BIG error 0xC0000040, ERRnomem 0x0008: Insufficient server memory to perform the requested operation."
	case strings.Contains(errorMessage, "STATUS_TOO_MANY_PAGING_FILES"):
			status = "Çok fazla sayfalama dosyası: STATUS_TOO_MANY_PAGING_FILES error 0xC0000097, ERRnomem 0x0008: Insufficient server memory to perform the requested operation."
	case strings.Contains(errorMessage, "STATUS_INSUFF_SERVER_RESOURCES"):
			status = "Sunucu kaynakları yetersiz: STATUS_INSUFF_SERVER_RESOURCES error 0xC0000205, ERRnomem 0x0008: Insufficient server memory to perform the requested operation."
	case strings.Contains(errorMessage, "STATUS_OS2_INVALID_ACCESS"):
			status = "Geçersiz erişim: STATUS_OS2_INVALID_ACCESS error 0x000C0001, ERRbadaccess 0x000C: Invalid open mode."
	case strings.Contains(errorMessage, "STATUS_ACCESS_DENIED"):
			status = "Erişim reddedildi: STATUS_ACCESS_DENIED error 0xC00000CA, ERRnomem 0x0008: Invalid open mode."
	case strings.Contains(errorMessage, "STATUS_DATA_ERROR"):
			status = "Veri hatası: STATUS_DATA_ERROR error 0xC000009C, ERRnomem 0x0008: Bad data. (May be generated by IOCTL calls on the server.)"
	case strings.Contains(errorMessage, "STATUS_DIRECTORY_NOT_EMPTY"):
			status = "Klasör boş değil: STATUS_DIRECTORY_NOT_EMPTY error 0xC0000101, ERRremcd 0x0010: Remove of directory failed because it was not empty."
	case strings.Contains(errorMessage, "STATUS_NOT_SAME_DEVICE"):
			status = "Aynı cihazda değil: STATUS_NOT_SAME_DEVICE error 0xC00000D4, ERRdiffdevice 0x0011: A file system operation (such as a rename) across two devices was attempted."
	case strings.Contains(errorMessage, "STATUS_NO_MORE_FILES"):
			status = "Daha fazla dosya yok: STATUS_NO_MORE_FILES error 0x80000006, ERRnofiles 0x0012: No (more) files found following a file search command."
	case strings.Contains(errorMessage, "STATUS_UNSUCCESSFUL"):
			status = "Başarısız: STATUS_UNSUCCESSFUL error 0xC0000001, ERRgeneral 0x001F: General error."
	case strings.Contains(errorMessage, "STATUS_SHARING_VIOLATION"):
			status = "Paylaşım ihlali: STATUS_SHARING_VIOLATION error 0xC0000043, ERRbadshare 0x0020: Sharing violation. A requested open mode conflicts with the sharing mode of an existing file handle."
	case strings.Contains(errorMessage, "STATUS_FILE_LOCK_CONFLICT"):
			status = "Dosya kilidi çakışması: STATUS_FILE_LOCK_CONFLICT error 0xC0000054, ERRlock 0x0021: A lock request specified an invalid locking mode, or conflicted with an existing file lock."
	case strings.Contains(errorMessage, "STATUS_LOCK_NOT_GRANTED"):
			status = "Kilit verilmedi: STATUS_LOCK_NOT_GRANTED error 0xC0000055, ERRlock 0x0021: A lock request specified an invalid locking mode, or conflicted with an existing file lock."
	case strings.Contains(errorMessage, "STATUS_END_OF_FILE"):
			status = "Dosya sonu: STATUS_END_OF_FILE error 0xC0000011, ERReof 0x0026: Attempted to read beyond the end of the file."
	case strings.Contains(errorMessage, "STATUS_NOT_SUPPORTED"):
			status = "Desteklenmiyor: STATUS_NOT_SUPPORTED error 0XC00000BB, ERRunsup 0x0032: This command is not supported by the server."
	case strings.Contains(errorMessage, "STATUS_OBJECT_NAME_COLLISION"):
			status = "Nesne adı çakışması: STATUS_OBJECT_NAME_COLLISION error 0xC0000035, ERRfilexists 0x0050: An attempt to create a file or directory failed because an object with the same pathname already exists."
	
			case strings.Contains(errorMessage, "STATUS_INVALID_PARAMETER"):
				status = "Geçersiz parametre: STATUS_INVALID_PARAMETER error 0xC000000D, ERRinvalidparam 0x0057 : A parameter supplied with the message is invalid."
			case strings.Contains(errorMessage, "STATUS_OS2_INVALID_LEVEL"):
				status = "Geçersiz bilgi seviyesi: STATUS_OS2_INVALID_LEVEL error 0x007C0001, ERRunknownlevel 0x007C : Invalid information level."
			case strings.Contains(errorMessage, "STATUS_OS2_NEGATIVE_SEEK"):
				status = "Negatif seek: STATUS_OS2_NEGATIVE_SEEK error 0x00830001, ERRinvalidseek 0x0083 : An attempt was made to seek to a negative absolute offset within a file."
			case strings.Contains(errorMessage, "STATUS_RANGE_NOT_LOCKED"):
				status = "Kilitli olmayan aralık: STATUS_RANGE_NOT_LOCKED error 0xC000007E, ERROR_NOT_LOCKED 0x009E : The byte range specified in an unlock request was not locked."
			case strings.Contains(errorMessage, "STATUS_OS2_NO_MORE_SIDS"):
				status = "Daha fazla SID yok: STATUS_OS2_NO_MORE_SIDS error 0x00710001, ERROR_NO_MORE_SEARCH_HANDLES 0x0071 : Maximum number of searches has been exhausted."
			case strings.Contains(errorMessage, "STATUS_OS2_CANCEL_VIOLATION"):
				status = "İptal ihlali: STATUS_OS2_CANCEL_VIOLATION error 0x00830001, ERROR_CANCEL_VIOLATION 0x00AD : No lock request was outstanding for the supplied cancel region."
			case strings.Contains(errorMessage, "STATUS_OS2_ATOMIC_LOCKS_NOT_SUPPORTED"):
				status = "Atomik kilitler desteklenmiyor: STATUS_OS2_ATOMIC_LOCKS_NOT_SUPPORTED error 0x00AE0001, ERROR_ATOMIC_LOCKS_NOT_SUPPORTED 0x00AE : The file system does not support atomic changes to the lock type."
			case strings.Contains(errorMessage, "STATUS_INVALID_INFO_CLASS"):
				status = "Geçersiz bilgi sınıfı: STATUS_INVALID_INFO_CLASS error 0xC0000003, ERRbadpipe 0x00E6 : Invalid named pipe."
			case strings.Contains(errorMessage, "STATUS_INVALID_PIPE_STATE"):
				status = "Geçersiz boru durumu: STATUS_INVALID_PIPE_STATE error 0xC00000AD, ERRbadpipe 0x00E6 : Invalid named pipe."
			case strings.Contains(errorMessage, "STATUS_INVALID_READ_MODE"):
				status = "Geçersiz okuma modu: STATUS_INVALID_READ_MODE error 0xC00000B4, ERRbadpipe 0x00E6 : Invalid named pipe."
			case strings.Contains(errorMessage, "STATUS_OS2_CANNOT_COPY"):
				status = "Kopyalama yapılamıyor: STATUS_OS2_CANNOT_COPY error 0x010A0001, ERROR_CANNOT_COPY 0x010A : The copy functions cannot be used."
			case strings.Contains(errorMessage, "STATUS_INSTANCE_NOT_AVAILABLE"):
				status = "Örnek mevcut değil: STATUS_INSTANCE_NOT_AVAILABLE error 0xC00000AB, ERRpipebusy 0x00E7 : All instances of the designated named pipe are busy."
			case strings.Contains(errorMessage, "STATUS_PIPE_NOT_AVAILABLE"):
				status = "Boru mevcut değil: STATUS_PIPE_NOT_AVAILABLE error 0xC00000AC, ERRpipebusy 0x00E7 : All instances of the designated named pipe are busy."
			case strings.Contains(errorMessage, "STATUS_PIPE_BUSY"):
				status = "Boru meşgul: STATUS_PIPE_BUSY error 0xC00000AE, ERRpipebusy 0x00E7 : All instances of the designated named pipe are busy."
			case strings.Contains(errorMessage, "STATUS_PIPE_CLOSING"):
				status = "Boru kapanıyor: STATUS_PIPE_CLOSING error 0xC00000B1, ERRpipeclosing 0x00E8: The designated named pipe is in the process of being closed."
			case strings.Contains(errorMessage, "STATUS_PIPE_EMPTY"):
				status = "Boru boş: STATUS_PIPE_EMPTY error 0xC00000D9, ERRpipeclosing 0x00E8: The designated named pipe is in the process of being closed."
			case strings.Contains(errorMessage, "STATUS_BUFFER_OVERFLOW"):
				status = "Tampon taşması: STATUS_BUFFER_OVERFLOW error 0x80000005, ERRmoredata 0x00EA: There is more data available to read on the designated named pipe."
			case strings.Contains(errorMessage, "STATUS_PIPE_DISCONNECTED"):
				status = "Boru bağlantısı kesildi: STATUS_PIPE_DISCONNECTED error 0xC00000B0, ERRnotconnected 0x00E9: The designated named pipe exists, but there is no server process listening on the server side."
			case strings.Contains(errorMessage, "STATUS_MORE_PROCESSING_REQUIRED"):
				status = "Daha fazla işlem yapılması gerekiyor: STATUS_MORE_PROCESSING_REQUIRED error 0xC0000016, ERRnotconnected 0x00E9: The designated named pipe exists, but there is no server process listening on the server side."
			case strings.Contains(errorMessage, "STATUS_EA_TOO_LARGE"):
				status = "EA çok büyük: STATUS_EA_TOO_LARGE error 0xC0000050, ERROR_EAS_DIDNT_FIT 0x0113: Either there are no extended attributes, or the available extended attributes did not fit into the response."
			case strings.Contains(errorMessage, "STATUS_OS2_EAS_DIDNT_FIT"):
				status = "EA sığmıyor: STATUS_OS2_EAS_DIDNT_FIT error 0x01130001, ERROR_EAS_DIDNT_FIT 0x0113: Either there are no extended attributes, or the available extended attributes did not fit into the response."
			case strings.Contains(errorMessage, "STATUS_EAS_NOT_SUPPORTED"):
				status = "EA desteklenmiyor: STATUS_EAS_NOT_SUPPORTED error 0xC000004F, ERROR_EAS_NOT_SUPPORTED 0x011A: The server file system does not support Extended Attributes."
			case strings.Contains(errorMessage, "STATUS_OS2_EA_ACCESS_DENIED"):
				status = "EA erişimi reddedildi: STATUS_OS2_EA_ACCESS_DENIED error 0x03E20001, ERROR_EA_ACCESS_DENIED 0x03E2: Access to the extended attribute was denied."
			case strings.Contains(errorMessage, "STATUS_NOTIFY_ENUM_DIR"):
				status = "Dizin değişiklikleri fazla: STATUS_NOTIFY_ENUM_DIR error 0x0000010C, ERR_NOTIFY_ENUM_DIR 0x03FE: More changes have occurred within the directory than will fit within the specified Change Notify response buffer."
			case strings.Contains(errorMessage, "STATUS_INVALID_SMB"):
				status = "Geçersiz SMB: STATUS_INVALID_SMB error 0x00010002, ERRerror 0x0001: Unspecified server error.<23>"
			case strings.Contains(errorMessage, "STATUS_WRONG_PASSWORD"):
				status = "Yanlış şifre: STATUS_WRONG_PASSWORD error 0xC000006A, ERRbadpw 0x0002: Invalid password."
			case strings.Contains(errorMessage, "STATUS_PATH_NOT_COVERED"):
				status = "Yol kapsanmıyor: STATUS_PATH_NOT_COVERED error 0xC0000257, ERRbadpath 0x0003: DFS pathname not on local server."
			case strings.Contains(errorMessage, "STATUS_NETWORK_ACCESS_DENIED"):
				status = "Ağ erişimi reddedildi: STATUS_NETWORK_ACCESS_DENIED error 0xC00000CA, ERRaccess 0x0004: Access denied. The specified UID does not have permission to execute the requested command within the current context (TID)."
			case strings.Contains(errorMessage, "STATUS_NETWORK_NAME_DELETED"):
				status = "Ağ adı silindi: STATUS_NETWORK_NAME_DELETED error 0xC00000C9, ERRinvtid 0x0005: The TID specified in the command was invalid. Earlier documentation, with the exception of [SNIA], refers to this error code as ERRinvnid (Invalid Network Path Identifier). [SNIA] uses both names.<24>"
			case strings.Contains(errorMessage, "STATUS_SMB_BAD_TID"):
				status = "Geçersiz TID: STATUS_SMB_BAD_TID error 0x00050002, ERRinvtid 0x0005: The TID specified in the command was invalid. Earlier documentation, with the exception of [SNIA], refers to this error code as ERRinvnid (Invalid Network Path Identifier). [SNIA] uses both names.<24>"
			case strings.Contains(errorMessage, "STATUS_BAD_NETWORK_NAME"):
				status = "Geçersiz ağ adı: STATUS_BAD_NETWORK_NAME error 0xC00000CC, ERRinvnetname 0x0006: Invalid server name in Tree Connect."
		
	
				case strings.Contains(errorMessage, "STATUS_BAD_DEVICE_TYPE"):
					status = " : STATUS_BAD_DEVICE_TYPE error 0xC00000CB, ERRinvdevice 0x0007: A printer request was made to a non-printer device or, conversely, a non-printer request was made to a printer device."
			
				case strings.Contains(errorMessage, "STATUS_SMB_BAD_COMMAND"):
					status = " : STATUS_SMB_BAD_COMMAND error 0x00160002, ERRbadcmd 0x0016: An unknown SMB command code was received by the server."
			
				case strings.Contains(errorMessage, "STATUS_PRINT_QUEUE_FULL"):
					status = " : STATUS_PRINT_QUEUE_FULL error 0xC00000C6, ERRqfull 0x0031: Print queue is full - too many queued items."
			
				case strings.Contains(errorMessage, "STATUS_NO_SPOOL_SPACE"):
					status = " : STATUS_NO_SPOOL_SPACE error 0xC00000C7, ERRqtoobig 0x0032: Print queue is full - no space for queued item, or queued item too big."
			
				case strings.Contains(errorMessage, "STATUS_PRINT_CANCELLED"):
					status = " : STATUS_PRINT_CANCELLED error 0xC00000C8, ERRinvpfid 0x0034: Invalid FID for print file."
			
				case strings.Contains(errorMessage, "STATUS_NOT_IMPLEMENTED"):
					status = " : STATUS_NOT_IMPLEMENTED error 0xC0000002, ERRsmbcmd 0x0040: Unrecognized SMB command code."
			
				case strings.Contains(errorMessage, "STATUS_UNEXPECTED_NETWORK_ERROR"):
					status = " : STATUS_UNEXPECTED_NETWORK_ERROR error 0xC00000C4, ERRsrverror 0x0041: Internal server error."
			
				case strings.Contains(errorMessage, "STATUS_NETWORK_ACCESS_DENIED"):
					status = " : STATUS_NETWORK_ACCESS_DENIED error 0xC00000CA, ERRbadpermits 0x0045: An invalid combination of access permissions for a file or directory was presented. The server cannot set the requested attributes."
			
				case strings.Contains(errorMessage, "STATUS_UNEXPECTED_NETWORK_ERROR"):
					status = " : STATUS_UNEXPECTED_NETWORK_ERROR error 0xC00000C4, ERRtimeout 0x0058: Operation timed out."
			
				case strings.Contains(errorMessage, "STATUS_IO_TIMEOUT"):
					status = " : STATUS_IO_TIMEOUT error 0xC00000B5, ERRtimeout 0x0058: Operation timed out."
			
				case strings.Contains(errorMessage, "STATUS_REQUEST_NOT_ACCEPTED"):
					status = " : STATUS_REQUEST_NOT_ACCEPTED error 0xC00000D0, ERRnoresource 0x0059: No resources currently available for this SMB request."
			
				case strings.Contains(errorMessage, "STATUS_TOO_MANY_SESSIONS"):
					status = " : STATUS_TOO_MANY_SESSIONS error 0xC00000CE, ERRtoomanyuids 0x005A: Too many UIDs active for this SMB connection."
			
				case strings.Contains(errorMessage, "STATUS_SMB_BAD_UID"):
					status = " : STATUS_SMB_BAD_UID error 0x005B0002, ERRbaduid 0x005B: The UID specified is not known as a valid ID on this server session."
			
				case strings.Contains(errorMessage, "STATUS_PIPE_DISCONNECTED"):
					status = " : STATUS_PIPE_DISCONNECTED error 0xC00000B0 , ERRnotconnected 0x00E9 : Write to a named pipe with no reader."
			
				case strings.Contains(errorMessage, "STATUS_SMB_USE_MPX"):
					status = " : STATUS_SMB_USE_MPX error 0x00FA0002 , ERRusempx 0x00FA : Temporarily unable to support RAW mode transfers. Use MPX mode."
			
				case strings.Contains(errorMessage, "STATUS_SMB_USE_STANDARD"):
					status = " : STATUS_SMB_USE_STANDARD error 0x00FB0002 , ERRusestd 0x00FB : Temporarily unable to support RAW or MPX mode transfers. Use standard read/write."
			
				case strings.Contains(errorMessage, "STATUS_SMB_CONTINUE_MPX"):
					status = " : STATUS_SMB_CONTINUE_MPX error 0x00FC0002 , ERRcontmpx 0x00FC : Continue in MPX mode. This error code is reserved for future use."
			
				case strings.Contains(errorMessage, "STATUS_ACCOUNT_DISABLED"):
					status = " : STATUS_ACCOUNT_DISABLED error 0xC0000072 , ERRaccountExpired 0x08BF : User account on the target machine is disabled or has expired."
			
				case strings.Contains(errorMessage, "STATUS_ACCOUNT_EXPIRED"):
					status = " : STATUS_ACCOUNT_EXPIRED error 0xC0000193 , ERRaccountExpired 0x08BF : User account on the target machine is disabled or has expired."
			
				case strings.Contains(errorMessage, "STATUS_INVALID_WORKSTATION"):
					status = " : STATUS_INVALID_WORKSTATION error 0xC0000070 , ERRbadClient 0x08C0 : The client does not have permission to access this server."
			
				case strings.Contains(errorMessage, "STATUS_INVALID_LOGON_HOURS"):
					status = " : STATUS_INVALID_LOGON_HOURS error 0xC000006F , ERRbadLogonTime 0x08C1 : Access to the server is not permitted at this time."
			
				case strings.Contains(errorMessage, "STATUS_PASSWORD_EXPIRED"):
					status = " : STATUS_PASSWORD_EXPIRED error 0xC0000071 , ERRpasswordExpired 0x08C2 : The user's password has expired."
			
				case strings.Contains(errorMessage, "STATUS_PASSWORD_MUST_CHANGE"):
					status = " : STATUS_PASSWORD_MUST_CHANGE error 0xC0000224 , ERRpasswordExpired 0x08C2 : The user's password has expired."
			
				case strings.Contains(errorMessage, "STATUS_SMB_NO_SUPPORT"):
					status = " : STATUS_SMB_NO_SUPPORT error 0xC0000224 , ERRnosupport 0xFFFF : Function not supported by the server."
			
				case strings.Contains(errorMessage, "STATUS_MEDIA_WRITE_PROTECTED"):
					status = " : STATUS_MEDIA_WRITE_PROTECTED error 0xC00000A2 , ERRnowrite 0x0013 : Attempt to modify a read-only file system."
			
				case strings.Contains(errorMessage, "STATUS_NO_MEDIA_IN_DEVICE"):
					status = " : STATUS_NO_MEDIA_IN_DEVICE error 0xC0000013 , ERRnotready 0x0015 : Drive not ready."
			
				case strings.Contains(errorMessage, "STATUS_INVALID_DEVICE_STATE"):
					status = " : STATUS_INVALID_DEVICE_STATE error 0xC0000184 , ERRbadcmd 0x0016 : Unknown command."
			
				case strings.Contains(errorMessage, "STATUS_DATA_ERROR"):
					status = " : STATUS_DATA_ERROR error 0xC000003E , ERRdata 0x0017 : Data error (incorrect CRC)."
			
				case strings.Contains(errorMessage, "STATUS_CRC_ERROR"):
					status = " : STATUS_CRC_ERROR error 0xC000003F , ERRdata 0x0017 : Data error (incorrect CRC)."
			
				case strings.Contains(errorMessage, "STATUS_DATA_ERROR"):
					status = " : STATUS_DATA_ERROR error 0xC000003E , ERRbadreq 0x0018 : Bad request structure length."
			
				case strings.Contains(errorMessage, "STATUS_DISK_CORRUPT_ERROR"):
					status = " : STATUS_DISK_CORRUPT_ERROR error 0xC0000032 , ERRbadmedia 0x001A : Unknown media type."
			
				case strings.Contains(errorMessage, "STATUS_NONEXISTENT_SECTOR"):
					status = " : STATUS_NONEXISTENT_SECTOR error 0xC0000015 , ERRbadsector 0x001B : Sector not found."
			
				case strings.Contains(errorMessage, "STATUS_DEVICE_PAPER_EMPTY"):
					status = " : STATUS_DEVICE_PAPER_EMPTY error 0x8000000E , ERRnopaper 0x001C : Printer out of paper."
			
				case strings.Contains(errorMessage, "STATUS_SHARING_VIOLATION"):
					status = " : STATUS_SHARING_VIOLATION error 0xC0000043 , ERRbadshare 0x0020 : An attempted open operation conflicts with an existing open."
			
				case strings.Contains(errorMessage, "STATUS_FILE_LOCK_CONFLICT"):
					status = " : STATUS_FILE_LOCK_CONFLICT error 0xC0000054 , ERRlock 0x0021 : A lock request specified an invalid locking mode, or conflicted with an existing file lock."
			
				case strings.Contains(errorMessage, "STATUS_WRONG_VOLUME"):
					status = " : STATUS_WRONG_VOLUME error 0xC0000012 , ERRwrongdisk 0x0022 : The wrong disk was found in a drive."
			
				case strings.Contains(errorMessage, "STATUS_DISK_FULL"):
					status = " : STATUS_DISK_FULL error 0xC000007F , ERRdiskfull 0x0027 : No space on file system."
    default:
        status = "Bilinmeyen hata"
    }
    Log445("aa hata %v", errorMessage)
    return status, hostname, fmt.Errorf(errorMessage)
}
defer s.Logoff()