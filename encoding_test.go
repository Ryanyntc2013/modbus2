package modbus

import (
	"testing"
)

func TestUint16ToBytes(t *testing.T) {
	var out []byte

	out = uint16ToBytes(BIG_ENDIAN, 0x4321)
	if len(out) != 2 {
		t.Errorf("expected 2 bytes, got %v", len(out))
	}
	if out[0] != 0x43 || out[1] != 0x21 {
		t.Errorf("expected {0x43, 0x21}, got {0x%02x, 0x%02x}", out[0], out[1])
	}

	out = uint16ToBytes(LITTLE_ENDIAN, 0x4321)
	if len(out) != 2 {
		t.Errorf("expected 2 bytes, got %v", len(out))
	}
	if out[0] != 0x21 || out[1] != 0x43 {
		t.Errorf("expected {0x21, 0x43}, got {0x%02x, 0x%02x}", out[0], out[1])
	}

	return
}

func TestUint16sToBytes(t *testing.T) {
	var out []byte

	out = uint16sToBytes(BIG_ENDIAN, []uint16{0x4321, 0x8765, 0xcba9})
	if len(out) != 6 {
		t.Errorf("expected 6 bytes, got %v", len(out))
	}
	if out[0] != 0x43 || out[1] != 0x21 {
		t.Errorf("expected {0x43, 0x21}, got {0x%02x, 0x%02x}", out[0], out[1])
	}
	if out[2] != 0x87 || out[3] != 0x65 {
		t.Errorf("expected {0x87, 0x65}, got {0x%02x, 0x%02x}", out[0], out[1])
	}
	if out[4] != 0xcb || out[5] != 0xa9 {
		t.Errorf("expected {0xcb, 0xa9}, got {0x%02x, 0x%02x}", out[0], out[1])
	}

	out = uint16sToBytes(LITTLE_ENDIAN, []uint16{0x4321, 0x8765, 0xcba9})
	if len(out) != 6 {
		t.Errorf("expected 6 bytes, got %v", len(out))
	}
	if out[0] != 0x21 || out[1] != 0x43 {
		t.Errorf("expected {0x21, 0x43}, got {0x%02x, 0x%02x}", out[0], out[1])
	}
	if out[2] != 0x65 || out[3] != 0x87 {
		t.Errorf("expected {0x65, 0x87}, got {0x%02x, 0x%02x}", out[0], out[1])
	}
	if out[4] != 0xa9 || out[5] != 0xcb {
		t.Errorf("expected {0xa9, 0xcb}, got {0x%02x, 0x%02x}", out[0], out[1])
	}

	return
}

func TestBytesToUint16(t *testing.T) {
	var result	uint16

	result	= bytesToUint16(BIG_ENDIAN, []byte{0x43, 0x21})
	if result != 0x4321 {
		t.Errorf("expected 0x4321, got 0x%04x", result)
	}

	result	= bytesToUint16(LITTLE_ENDIAN, []byte{0x43, 0x21})
	if result != 0x2143 {
		t.Errorf("expected 0x2143, got 0x%04x", result)
	}

	return
}

func TestBytesToUint16s(t *testing.T) {
	var results	[]uint16

	results	= bytesToUint16s(BIG_ENDIAN, []byte{0x11, 0x22, 0x33, 0x44})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 0x1122 {
		t.Errorf("expected 0x1122, got 0x%04x", results[0])
	}
	if results[1] != 0x3344 {
		t.Errorf("expected 0x3344, got 0x%04x", results[1])
	}

	results	= bytesToUint16s(LITTLE_ENDIAN, []byte{0x11, 0x22, 0x33, 0x44})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 0x2211 {
		t.Errorf("expected 0x2211, got 0x%04x", results[0])
	}
	if results[1] != 0x4433 {
		t.Errorf("expected 0x4433, got 0x%04x", results[1])
	}

	return
}

func TestUint32ToBytes(t *testing.T) {
	var out []byte

	out = uint32ToBytes(BIG_ENDIAN, HIGH_WORD_FIRST, 0x87654321)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0x87 || out[1] != 0x65 || // first word
	   out[2] != 0x43 || out[3] != 0x21 {  // second word
		t.Errorf("expected {0x87, 0x65, 0x43, 0x21}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	out = uint32ToBytes(BIG_ENDIAN, LOW_WORD_FIRST, 0x87654321)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0x43 || out[1] != 0x21 || out[2] != 0x87 || out[3] != 0x65 {
		t.Errorf("expected {0x43, 0x21, 0x87, 0x65}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	out = uint32ToBytes(LITTLE_ENDIAN, LOW_WORD_FIRST, 0x87654321)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0x21 || out[1] != 0x43 || out[2] != 0x65 || out[3] != 0x87 {
		t.Errorf("expected {0x21, 0x43, 0x65, 0x87}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	out = uint32ToBytes(LITTLE_ENDIAN, HIGH_WORD_FIRST, 0x87654321)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0x65 || out[1] != 0x87 || out[2] != 0x21 || out[3] != 0x43 {
		t.Errorf("expected {0x65, 0x87, 0x21, 0x43}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	return
}

func TestBytesToUint32s(t *testing.T) {
	var results []uint32

	results	= bytesToUint32s(BIG_ENDIAN, HIGH_WORD_FIRST, []byte{
		0x87, 0x65, 0x43, 0x21,
		0x00, 0x11, 0x22, 0x33,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 0x87654321 {
		t.Errorf("expected 0x87654321, got 0x%08x", results[0])
	}
	if results[1] != 0x00112233 {
		t.Errorf("expected 0x00112233, got 0x%08x", results[1])
	}

	results	= bytesToUint32s(BIG_ENDIAN, LOW_WORD_FIRST, []byte{
		0x87, 0x65, 0x43, 0x21,
		0x00, 0x11, 0x22, 0x33,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 0x43218765 {
		t.Errorf("expected 0x43218765, got 0x%08x", results[0])
	}
	if results[1] != 0x22330011 {
		t.Errorf("expected 0x22330011, got 0x%08x", results[1])
	}

	results	= bytesToUint32s(LITTLE_ENDIAN, LOW_WORD_FIRST, []byte{
		0x87, 0x65, 0x43, 0x21,
		0x00, 0x11, 0x22, 0x33,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 0x21436587 {
		t.Errorf("expected 0x21436587, got 0x%08x", results[0])
	}
	if results[1] != 0x33221100 {
		t.Errorf("expected 0x33221100, got 0x%08x", results[1])
	}

	results	= bytesToUint32s(LITTLE_ENDIAN, HIGH_WORD_FIRST, []byte{
		0x87, 0x65, 0x43, 0x21,
		0x00, 0x11, 0x22, 0x33,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 0x65872143 {
		t.Errorf("expected 0x65872143, got 0x%08x", results[0])
	}
	if results[1] != 0x11003322 {
		t.Errorf("expected 0x11003322, got 0x%08x", results[1])
	}

	return
}

func TestFloat32ToBytes(t *testing.T) {
	var out	[]byte

	out = float32ToBytes(BIG_ENDIAN, HIGH_WORD_FIRST, 1.234)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0x3f || out[1] != 0x9d || out[2] != 0xf3 || out[3] != 0xb6 {
		t.Errorf("expected {0x3f, 0x9d, 0xf3, 0xb6}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	out = float32ToBytes(BIG_ENDIAN, LOW_WORD_FIRST, 1.234)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0xf3 || out[1] != 0xb6 || out[2] != 0x3f || out[3] != 0x9d {
		t.Errorf("expected {0xf3, 0xb6, 0x3f, 0x9d}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	out = float32ToBytes(LITTLE_ENDIAN, LOW_WORD_FIRST, 1.234)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0xb6 || out[1] != 0xf3 || out[2] != 0x9d || out[3] != 0x3f {
		t.Errorf("expected {0xb6, 0xf3, 0x9d, 0x3f}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	out = float32ToBytes(LITTLE_ENDIAN, HIGH_WORD_FIRST, 1.234)
	if len(out) != 4 {
		t.Errorf("expected 4 bytes, got %v", len(out))
	}
	if out[0] != 0x9d || out[1] != 0x3f || out[2] != 0xb6 || out[3] != 0xf3 {
		t.Errorf("expected {0x9d, 0x3f, 0xb6, 0xf3}, got {0x%02x, 0x%02x, 0x%02x, 0x%02x}",
			 out[0], out[1], out[2], out[3])
	}

	return
}

func TestBytesToFloat32s(t *testing.T) {
	var results []float32

	results	= bytesToFloat32s(BIG_ENDIAN, HIGH_WORD_FIRST, []byte{
		0x3f, 0x9d, 0xf3, 0xb6,
		0x40, 0x49, 0x0f, 0xdb,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 1.234 {
		t.Errorf("expected 1.234, got %.04f", results[0])
	}
	if results[1] != 3.14159274101 {
		t.Errorf("expected 3.14159274101, got %.09f", results[1])
	}

	results	= bytesToFloat32s(BIG_ENDIAN, LOW_WORD_FIRST, []byte{
		0xf3, 0xb6, 0x3f, 0x9d,
		0x0f, 0xdb, 0x40, 0x49,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 1.234 {
		t.Errorf("expected 1.234, got %.04f", results[0])
	}
	if results[1] != 3.14159274101 {
		t.Errorf("expected 3.14159274101, got %.09f", results[1])
	}

	results	= bytesToFloat32s(LITTLE_ENDIAN, LOW_WORD_FIRST, []byte{
		0xb6, 0xf3, 0x9d, 0x3f,
		0xdb, 0x0f, 0x49, 0x40,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 1.234 {
		t.Errorf("expected 1.234, got %.04f", results[0])
	}
	if results[1] != 3.14159274101 {
		t.Errorf("expected 3.14159274101, got %.09f", results[1])
	}

	results	= bytesToFloat32s(LITTLE_ENDIAN, HIGH_WORD_FIRST, []byte{
		0x9d, 0x3f, 0xb6, 0xf3,
		0x49, 0x40, 0xdb, 0x0f,
	})
	if len(results) != 2 {
		t.Errorf("expected 2 values, got %v", len(results))
	}
	if results[0] != 1.234 {
		t.Errorf("expected 1.234, got %.04f", results[0])
	}
	if results[1] != 3.14159274101 {
		t.Errorf("expected 3.14159274101, got %.09f", results[1])
	}

	return
}
