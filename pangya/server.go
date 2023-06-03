// Copyright (C) 2018-2023, John Chadwick <john@jchw.io>
//
// Permission to use, copy, modify, and/or distribute this software for any purpose
// with or without fee is hereby granted, provided that the above copyright notice
// and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
// REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND
// FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
// INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM LOSS
// OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR OTHER
// TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE OF
// THIS SOFTWARE.
//
// SPDX-FileCopyrightText: Copyright (c) 2018-2023 John Chadwick
// SPDX-License-Identifier: ISC

package pangya

type ChannelEntry struct {
	ChannelName string `struct:"[64]byte"`
	MaxUsers    uint16
	NumUsers    uint16
	Unknown1    uint16
	Unknown2    uint16
	Unknown3    [5]byte
}

// ServerEntry represents a server in a ServerListMessage.
type ServerEntry struct {
	ServerName string `struct:"[40]byte"`
	ServerID   uint32
	MaxUsers   uint32
	NumUsers   uint32
	IPAddress  string `struct:"[18]byte"`
	Port       uint16
	Unknown3   uint16
	Flags      uint16
	Unknown4   [16]byte

	Count    byte `struct:"sizeof=Channels"`
	Channels []ChannelEntry
}