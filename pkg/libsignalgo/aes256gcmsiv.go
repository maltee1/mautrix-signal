// mautrix-signal - A Matrix-signal puppeting bridge.
// Copyright (C) 2023 Sumner Evans
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package libsignalgo

/*
#cgo LDFLAGS: -lsignal_ffi -ldl -lm
#include "./libsignal-ffi.h"
*/
import "C"
import "runtime"

type AES256_GCM_SIV struct {
	ptr *C.SignalAes256GcmSiv
}

func wrapAES256_GCM_SIV(ptr *C.SignalAes256GcmSiv) *AES256_GCM_SIV {
	aes := &AES256_GCM_SIV{ptr: ptr}
	runtime.SetFinalizer(aes, (*AES256_GCM_SIV).Destroy)
	return aes
}

func NewAES256_GCM_SIV(key []byte) (*AES256_GCM_SIV, error) {
	var aes *C.SignalAes256GcmSiv
	signalFfiError := C.signal_aes256_gcm_siv_new(&aes, BytesToBuffer(key))
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return wrapAES256_GCM_SIV(aes), nil
}

func (aes *AES256_GCM_SIV) Destroy() error {
	runtime.SetFinalizer(aes, nil)
	return wrapError(C.signal_aes256_gcm_siv_destroy(aes.ptr))
}

func (aes *AES256_GCM_SIV) Encrypt(plaintext, nonce, associatedData []byte) ([]byte, error) {
	var encrypted C.SignalOwnedBuffer = C.SignalOwnedBuffer{}

	signalFfiError := C.signal_aes256_gcm_siv_encrypt(&encrypted, aes.ptr, BytesToBuffer(plaintext), BytesToBuffer(nonce), BytesToBuffer(associatedData))
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return CopySignalOwnedBufferToBytes(encrypted), nil
}

func (aes *AES256_GCM_SIV) Decrypt(ciphertext, nonce, associatedData []byte) ([]byte, error) {
	var decrypted C.SignalOwnedBuffer = C.SignalOwnedBuffer{}
	signalFfiError := C.signal_aes256_gcm_siv_decrypt(&decrypted, aes.ptr, BytesToBuffer(ciphertext), BytesToBuffer(nonce), BytesToBuffer(associatedData))
	if signalFfiError != nil {
		return nil, wrapError(signalFfiError)
	}
	return CopySignalOwnedBufferToBytes(decrypted), nil
}
