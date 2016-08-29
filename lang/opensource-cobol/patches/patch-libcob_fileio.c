$NetBSD$

Upstream patch: https://github.com/opensourcecobol/opensource-cobol/pull/31

--- libcob/fileio.c.orig	2016-08-29 04:16:28.000000000 +0000
+++ libcob/fileio.c
@@ -215,6 +215,7 @@ struct dirent		*listdir_filedata;
 #define        READOPTSSIZE  4
 #define        STARTCONDSIZE 2
 #define        EXCPTCODESIZE 6
+#define        FNSTATUSSIZE  3
 
 cob_file		*cob_error_file;
 
@@ -4120,7 +4121,7 @@ cob_invoke_fun (int operate, char *f, co
 	char	oper[OPENMODESIZE];
 	char	excpcode[EXCPTCODESIZE];
 	char	*p_excpcode = excpcode;
-	char	tmpfnstatus[2];
+	char	tmpfnstatus[FNSTATUSSIZE];
 	char	*p_tmpfnstatus = tmpfnstatus;
 	int		status1 = 0;
 	int	(*funcint)();
