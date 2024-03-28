package caj

/*
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Cross-platform support for Windows & Linux.
#if !(defined(CAJ2PDF_OS_WINDOWS) || defined(CAJ2PDF_OS_LINUX))
#if (defined(_WIN32) || defined(_WIN64) || defined(__WIN32__) || defined(__TOS_WIN__) || defined(__WINDOWS__))
#define CAJ2PDF_OS_WINDOWS
#endif
#if (defined(linux) || defined(__linux) || defined(__linux__) || defined(__gnu_linux__))
#define CAJ2PDF_OS_LINUX
#endif
#endif
#if !(defined(CAJ2PDF_OS_WINDOWS) || defined(CAJ2PDF_OS_LINUX))
#error unsupported platform
#endif
#if defined(CAJ2PDF_OS_WINDOWS) && defined(CAJ2PDF_OS_LINUX)
#error conflicting platform definition
#endif

#if defined(CAJ2PDF_OS_WINDOWS)
#define WIN32_LEAN_AND_MEAN
#include <windows.h>
#elif defined(CAJ2PDF_OS_LINUX)
#include <dlfcn.h>
#endif

struct Parameter
{
	size_t cb;
	uint32_t flag[2];
	char *src;
	char *extname;
	// Function pointers for open, read, seek, tell, eof, and close.
	void *pfnFILE[6];
	char *dest;
	// Function pointers.
	void *pfnoss[4];
};

#if defined(CAJ2PDF_OS_LINUX)
// CAJFILE_Init required on Linux.
typedef int (*PCAJFILE_Init)(char const *exeDir, char const **, int, char const *tmpDir);
#endif
typedef int (*PCAJFILE_DistillPageEx1)(struct Parameter *param);

static char src[4096], extname[256], dest[4096];

//export cajConvertC
int cajConvertC(const char *sourcePathname, const char *sourceExtname, const char *destinationPathname)
{
#if defined(CAJ2PDF_OS_WINDOWS)
#ifdef _WIN64
	HMODULE handle = LoadLibrary(TEXT(".\\ReaderEx_x64.dll"));
#else
	HMODULE handle = LoadLibrary(TEXT(".\\ReaderEx.dll"));
#endif
#elif defined(CAJ2PDF_OS_LINUX)
	void *handle = dlopen("./libreaderex_x64.so", RTLD_LAZY);
#endif
	if (handle)
	{
		printf("[D] handle = %p\n", handle);
#if defined(CAJ2PDF_OS_WINDOWS)
		PCAJFILE_DistillPageEx1 CAJFILE_DistillPageEx1;
		// ReaderEx_DIS_C 2.3.0 Build 3982: ordinal value = 216
		// 2.0.3949.0 to 2.3.3983.0, 2.3.3982.0 x64 : ordinal value = 216
		//
		// The customized jbig introduced between 2.0.3920.0 and 2.0.3949.0
		// ImageCodec.dll jumped from 800k to 930k
		//
		// 2.0.3920.0                            HN : jpeg + jbig2
		// 2.0.3684.0, 2.0.3799.0                HN : jpeg only
		*(void **)(&CAJFILE_DistillPageEx1) = GetProcAddress(handle, (LPCSTR)216);
#elif defined(CAJ2PDF_OS_LINUX)
		// CAJFILE_Init required on Linux.
		PCAJFILE_Init CAJFILE_Init;
		PCAJFILE_DistillPageEx1 CAJFILE_DistillPageEx1;
		*(void **)(&CAJFILE_Init) = dlsym(handle, "CAJFILE_Init");
		if (CAJFILE_Init)
		{
			printf("[D] CAJFILE_Init = %p\n", CAJFILE_Init);
			int result = CAJFILE_Init(".", NULL, 0, "./tmp");
			if (result)
			{
				printf("[E] CAJFILE_Init failed (%d)\n", result);
				goto cleanup;
			}
		}
		else
		{
			printf("[E] dlsym failed (%s)\n", dlerror());
			goto cleanup;
		}
		*(void **)(&CAJFILE_DistillPageEx1) = dlsym(handle, "CAJFILE_DistillPageEx1");
#endif
		if (CAJFILE_DistillPageEx1)
		{
			printf("[D] CAJFILE_DistillPageEx1 = %p\n", CAJFILE_DistillPageEx1);
			struct Parameter param;
			memset(&param, 0, sizeof param);
			param.cb = sizeof param;
			param.flag[1] = 0x26;

			strncpy(src, sourcePathname, sizeof(src) - 1);
			src[sizeof(src) - 1] = '\0';
			strncpy(extname, sourceExtname, sizeof(extname) - 1);
			extname[sizeof(extname) - 1] = '\0';
			strncpy(dest, destinationPathname, sizeof(dest) - 1);
			dest[sizeof(dest) - 1] = '\0';

			param.src = src;
			param.extname = extname;
			param.dest = dest;

			int result = CAJFILE_DistillPageEx1(&param);
			if (result)
			{
				printf("[I] Success\n");
			}
			else
			{
				printf("[E] CAJFILE_DistillPageEx1 failed (%d)\n", result);
				goto cleanup;
			}
		}
		else
		{
#if defined(CAJ2PDF_OS_WINDOWS)
			printf("[E] GetProcAddress failed (%lu)\n", GetLastError());
#elif defined(CAJ2PDF_OS_LINUX)
			printf("[E] dlsym failed (%s)\n", dlerror());
#endif
			goto cleanup;
		}
	cleanup:
#if defined(CAJ2PDF_OS_WINDOWS)
		// If the function succeeds, the return value is nonzero.
		// If the function fails, the return value is zero.
		if (!FreeLibrary(handle))
		{
			printf("[E] FreeLibrary failed (%lu)\n", GetLastError());
		}
#elif defined(CAJ2PDF_OS_LINUX)
		// On success, dlclose() returns 0; on error, it returns a nonzero value.
		if (dlclose(handle))
		{
			printf("[E] dlclose failed (%s)\n", dlerror());
		}
#endif
	}
	else
	{
#if defined(CAJ2PDF_OS_WINDOWS)
		printf("[E] LoadLibrary failed (%lu)\n", GetLastError());
#elif defined(CAJ2PDF_OS_LINUX)
		printf("[E] dlopen failed (%s)\n", dlerror());
#endif
	}
}

*/
import "C"
import "unsafe"

func Caj2pdf(source string, des string) {
	sourcePathname := C.CString(source)
	sourceExtname := C.CString(".caj")
	destinationPathname := C.CString(des)
	defer C.free(unsafe.Pointer(sourcePathname))
	defer C.free(unsafe.Pointer(sourceExtname))
	defer C.free(unsafe.Pointer(destinationPathname))
	C.cajConvertC(sourcePathname, sourceExtname, destinationPathname)
}
