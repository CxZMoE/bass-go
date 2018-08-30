package bass

// #cgo LDFLAGS: -Llib -lbass
// #include "head/bass.h"
import "C"
import (
	"unsafe"
)

const ( //Define some constant of bass
	BASS_ACTIVE_STOPPED = C.BASS_ACTIVE_STOPPED
	BASS_ACTIVE_PLAYING = C.BASS_ACTIVE_PLAYING
	BASS_ACTIVE_PAUSED  = C.BASS_ACTIVE_PAUSED
	BASS_ACTIVE_STALLED = C.BASS_ACTIVE_STALLED
)

type ulong C.ulong

//------Initialization,into,tec..
func Bass_Init() C.int { //Initializes an output device.
	return C.BASS_Init(-1, 44100, 0, nil, nil)
}

func Bass_Free() C.int { //Frees all resources used by the output device, including all its samples, streams and MOD musics.
	return C.BASS_Free()
}

func Bass_GetVersion() C.ulong { //Retrieves the version of BASS that is loaded.
	return C.BASS_GetVersion()
}

func Bass_GetVolume() C.float { //Retrieves the current master volume level.
	return C.BASS_GetVolume()
}

func Bass_GetInfo(info *C.BASS_INFO) C.int { //Retrieves information on the device being used.
	return C.BASS_GetInfo(info)
}

func Bass_Pause() C.int { //Stops the output, pausing all musics/samples/streams on it.
	return C.BASS_Pause()
}

func Bass_SetDevice(device C.ulong) C.int { //Sets the device to use for subsequent calls in the current thread.
	return C.BASS_SetDevice(device)
}

func Bass_GetDevice(device C.ulong) C.ulong { //Retrieves the device setting of the current thread.
	return C.BASS_GetDevice()
}

func Bass_GetCPU() C.float { //Retrieves the current CPU usage of BASS.
	return C.BASS_GetCPU()
}

//--------------------------------------

//------Streams-------------------------
func Bass_StreamCreate(freq C.ulong, proc *C.STREAMPROC, user unsafe.Pointer) C.ulong { //Creates a user sample stream.
	return C.BASS_StreamCreate(freq, 2, C.BASS_SAMPLE_FLOAT, proc, user)
}
func Bass_StreamCreateFile(mem C.int, file string, offset C.ulonglong, length C.ulonglong) C.ulong { //Creates a sample stream from an MP3, MP2, MP1, OGG, WAV, AIFF or plugin supported file.
	return C.BASS_StreamCreateFile(mem, unsafe.Pointer(C.CString(file)), offset, length, C.BASS_SAMPLE_FLOAT)
}
func Bass_StreamCreateURL(url string, offset C.ulong, flags C.ulong, proc *C.DOWNLOADPROC, user unsafe.Pointer) C.ulong { //ates a sample stream from an MP3, MP2, MP1, OGG, WAV, AIFF or plugin supported file on the internet, optionally receiving the downloaded data in a callback function.
	return C.BASS_StreamCreateURL(C.CString(url), offset, flags, proc, user)
}

func Bass_StreamPutData(handle C.ulong, buffer []byte, length int) uint32 {

	return uint32(C.BASS_StreamPutData(handle, C.CBytes(buffer), C.ulong(length)))
}

//--------------------------------------

//------Channels-------------------------
func Bass_ChannelPlay(handle C.ulong, restart C.int) C.int { //Starts (or resumes) playback of a sample, stream, MOD music, or recording.
	return C.BASS_ChannelPlay(handle, restart)
}

func BASS_ChannelPause(handle C.ulong) C.int { //Pauses a sample, stream, MOD music, or recording.
	return C.BASS_ChannelPause(handle)
}

func BASS_ChannelStop(handle C.ulong) C.int { //Stops a sample, stream, MOD music, or recording.
	return C.BASS_ChannelStop(handle)
}

func BASS_ChannelBytes2Seconds(handle C.ulong, pos C.ulonglong) C.double { //Translates a byte position into time (seconds), based on a channel's format.
	return C.BASS_ChannelBytes2Seconds(handle, pos)
}

func BASS_ChannelSeconds2Bytes(handle C.ulong, pos C.double) C.ulonglong { //Translates a time (seconds) position into bytes, based on a channel's format.
	return BASS_ChannelSeconds2Bytes(handle, pos)
}

func BASS_ChannelIsActive(handle C.ulong) C.ulong { //Checks if a sample, stream, or MOD music is active (playing) or stalled. Can also check if a recording is in progress.
	return C.BASS_ChannelIsActive(handle)
}

func BASS_ChannelGetPosition(handle C.ulong, mode C.ulong) C.ulonglong { //Retrieves the playback position of a sample, stream, or MOD music. Can also be used with a recording channel.
	return C.BASS_ChannelGetPosition(handle, mode)
}

func BASS_ChannelSetPosition(handle C.ulong, pos C.ulonglong, mode C.ulong) C.int { //Sets the playback position of a sample, MOD music, or stream.
	return BASS_ChannelSetPosition(handle, pos, mode)
}

func BASS_ChannelSetAttribute(handle C.ulong, attrib C.ulong, value C.float) C.int { //Sets the value of a channel's attribute.
	return BASS_ChannelSetAttribute(handle, attrib, value)
}

func BASS_ChannelUpdate(handle C.ulong, length C.ulong) C.int { //Updates the playback buffer of a stream or MOD music.
	return BASS_ChannelUpdate(handle, length)
}

//---------------------------------------
